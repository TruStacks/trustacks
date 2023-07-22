package plan

import (
	"encoding/json"
	"errors"
	"fmt"

	"dagger.io/dagger"
	"github.com/charmbracelet/log"
	mapset "github.com/deckarep/golang-set/v2"
)

type ActionPlan struct {
	Actions     []string               `json:"actions"`
	InputFields map[string]interface{} `json:"inputs"`
	id          string
	artifacts   *ArtifactStore
}

type Action struct {
	Name            string
	Image           string
	Script          func(*dagger.Container, map[string]interface{}, *ActionUtilities) error
	State           State
	InputArtifacts  []Artifact
	OutputArtifacts []Artifact
	Caches          []string
	Secrets         []string
}

func (ap *ActionPlan) Schedule() (map[State]mapset.Set[*Action], error) {
	schedule := map[State]mapset.Set[*Action]{}
	inputs := map[Artifact]mapset.Set[State]{}
	for _, actionName := range ap.Actions {
		for registeredActionName, registeredAction := range registeredActions {
			if actionName == registeredActionName {
				for _, artifact := range registeredAction.InputArtifacts {
					if _, ok := inputs[artifact]; !ok {
						inputs[artifact] = mapset.NewSet[State]()
					}
					inputs[artifact].Add(registeredAction.State)
				}
				if _, ok := schedule[registeredAction.State]; !ok {
					schedule[registeredAction.State] = mapset.NewSet[*Action]()
				}
				schedule[registeredAction.State].Add(registeredAction)
			}
		}
	}
	unmatchableWatchList := mapset.NewSet[*Action]()
	for {
		if schedule[OnDemandState] == nil || len(schedule[OnDemandState].ToSlice()) == 0 {
			break
		}
		for _, action := range schedule[OnDemandState].ToSlice() {
			inputActionMatched := false
			for _, artifact := range action.OutputArtifacts {
				if states, ok := inputs[artifact]; ok {
					inputActionMatched = true
					states := states.ToSlice()
					firstStateOccurance := states[0]
					for i := 1; i < len(states); i++ {
						if states[i] > firstStateOccurance {
							firstStateOccurance = states[i]
						}
					}
					if firstStateOccurance == OnDemandState {
						if unmatchableWatchList.Contains(action) {
							return nil, fmt.Errorf("outputs for on-demand action '%s' cannot be resolved", action.Name)
						}
						unmatchableWatchList.Add(action)
						continue
					}
					for _, artifact := range action.InputArtifacts {
						inputs[artifact].Add(firstStateOccurance)
						inputs[artifact].Remove(OnDemandState)
					}
					schedule[firstStateOccurance].Add(action)
					schedule[OnDemandState].Remove(action)
				}
			}
			if !inputActionMatched {
				err := fmt.Errorf("one or more outputs from on-demand action '%s' could not be matched to an action in an activity state. this action will be excluded from the action plan", action.Name)
				log.Warn("", "err", err)
				schedule[OnDemandState].Remove(action)
			}
		}
	}
	return schedule, nil
}

func (ap *ActionPlan) AddAction(name string, InputFields []string) {
	ap.Actions = append(ap.Actions, name)
	for _, input := range InputFields {
		ap.InputFields[input] = nil
	}
}

func (ap *ActionPlan) ToJson() (string, error) {
	data, err := json.Marshal(ap)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

type deferAction bool

func (ap *ActionPlan) runAction(source string, action *Action, client *dagger.Client) (deferAction, error) {
	for _, artifact := range action.InputArtifacts {
		if _, ok := ap.artifacts.artifacts[artifact]; !ok {
			return true, nil
		}
	}
	log.Info("Queuing action", "action", action.Name)
	container := client.Container().From(action.Image)
	container = container.WithMountedDirectory("/src", client.Host().Directory(source)).WithWorkdir("/src")
	for _, path := range action.Caches {
		container = container.WithMountedCache(path, client.CacheVolume(ap.id+path))
	}
	return false, action.Script(container, ap.InputFields, newActionUtilities(client, ap.artifacts))
}

// NewActionPlan .
func NewActionPlan(client *dagger.Client) *ActionPlan {
	plan := &ActionPlan{
		InputFields: make(map[string]interface{}),
	}
	if client != nil {
		plan.artifacts = newArtifactStore(client)
	}
	return plan
}

// Run .
func Run(source, spec string, client *dagger.Client, states []State) error {
	ap := NewActionPlan(client)
	defer func() {
		for _, mount := range ap.artifacts.mounts {
			mount.Close()
		}
	}()
	if err := json.Unmarshal([]byte(spec), &ap); err != nil {
		return err
	}
	for _, input := range ap.InputFields {
		if input == nil {
			return errors.New("action plan cannot contain null user-defined inputs")
		}
	}
	schedule, err := ap.Schedule()
	if err != nil {
		return err
	}
	for _, actions := range schedule {
		for action := range actions.Iter() {
			log.Info(fmt.Sprintf("> %s", action.Name))
		}
	}
	for _, state := range states {
		if actions, ok := schedule[state]; ok {
			for {
				if len(actions.ToSlice()) == 0 {
					break
				}
				for _, action := range actions.ToSlice() {
					deferred, err := ap.runAction(source, action, client)
					if err != nil {
						return err
					}
					if !deferred {
						actions.Remove(action)
					}
				}
			}
		}
	}
	return nil
}

// actions .
var registeredActions = map[string]*Action{}

// RegisterAction .
func RegisterAction(action *Action) {
	registeredActions[action.Name] = action
}
