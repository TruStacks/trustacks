package plan

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"dagger.io/dagger"
	"github.com/charmbracelet/log"
	mapset "github.com/deckarep/golang-set/v2"
)

type ActionSpec struct {
	Name        string `json:"name"`
	DisplayName string `json:"displayName,omitempty"`
	Description string `json:"description,omitempty"`
}

type ActionPlan struct {
	Actions   []ActionSpec           `json:"actions"`
	Fields    []string               `json:"fields"`
	Inputs    map[string]interface{} `json:"inputs,omitempty"`
	id        string
	artifacts *ArtifactStore
}

type Action struct {
	Name                   string
	Image                  string
	Script                 func(*dagger.Container, map[string]interface{}, *ActionUtilities) error
	Stage                  Stage
	InputArtifacts         []Artifact
	OptionalInputArtifacts []Artifact
	OutputArtifacts        []Artifact
	Caches                 []string
	Secrets                []string
}

type scheduler struct {
	requiredInputs map[Artifact]mapset.Set[Stage]
	optionalInputs map[Artifact]mapset.Set[Stage]
}

func newScheduler() *scheduler {
	return &scheduler{
		requiredInputs: map[Artifact]mapset.Set[Stage]{},
		optionalInputs: map[Artifact]mapset.Set[Stage]{},
	}
}

func (s *scheduler) assignActivityStage(actions []string) map[Stage]mapset.Set[*Action] {
	assignments := map[Stage]mapset.Set[*Action]{}
	for _, actionName := range actions {
		for registeredActionName, registeredAction := range registeredActions {
			if actionName == registeredActionName {
				if _, ok := assignments[registeredAction.Stage]; !ok {
					assignments[registeredAction.Stage] = mapset.NewSet[*Action]()
				}
				assignments[registeredAction.Stage].Add(registeredAction)
			}
		}
	}
	return assignments
}

func (s *scheduler) bindActionInputs(assignments map[Stage]mapset.Set[*Action]) {
	for stage, actions := range assignments {
		for _, action := range actions.ToSlice() {
			for _, artifact := range action.InputArtifacts {
				if _, ok := s.requiredInputs[artifact]; !ok {
					s.requiredInputs[artifact] = mapset.NewSet[Stage]()
				}
				s.requiredInputs[artifact].Add(stage)
			}
			for _, artifact := range action.OptionalInputArtifacts {
				if _, ok := s.optionalInputs[artifact]; !ok {
					s.optionalInputs[artifact] = mapset.NewSet[Stage]()
				}
				s.optionalInputs[artifact].Add(stage)
			}
		}
	}
}

func (s *scheduler) assignOnDemandActions(assignments map[Stage]mapset.Set[*Action]) error {
	unmatchableWatchList := mapset.NewSet[*Action]()
	for {
		// if on demand actions are nil or empty then all actions are assigned
		// or no on demand actions exist in the schedule.
		if assignments[OnDemandStage] == nil || assignments[OnDemandStage].Cardinality() == 0 {
			break
		}
		for _, action := range assignments[OnDemandStage].ToSlice() {
			matched := false
			// check if the on demand action output artifacts have actions
			// in an activity stage that require their outputs.
			for _, artifact := range action.OutputArtifacts {
				if stages, ok := s.requiredInputs[artifact]; ok {
					matched = true
					stages := stages.ToSlice()
					firstOccurance := stages[0]
					for i := 1; i < len(stages); i++ {
						if stages[i] < firstOccurance {
							firstOccurance = stages[i]
						}
					}
					// if the first occurance of the input is in the on demand
					// stage, then we must wait until the next loop iteration
					// for an input transfer.
					if firstOccurance == OnDemandStage {
						// if the action is already in the unmatched list then
						// the matching failed for two loop cycles.
						// This means that the associated input could not be
						// tranferred to to a finite activity stage and cannot
						// be resolved.
						if unmatchableWatchList.Contains(action) {
							return fmt.Errorf("outputs for on-demand action '%s' cannot be resolved", action.Name)
						}
						// add the action to the watch list until the next loop
						// cycle.
						unmatchableWatchList.Add(action)
						continue
					}
					// transfer the action from the on demand staage to the
					// required activity stage.
					assignments[firstOccurance].Add(action)
					assignments[OnDemandStage].Remove(action)
					// it is possible for an on demand action to have inputs.
					// transfer inputs along with the action to the appropriate
					// stage and remove those inputs from the on demand stage
					// since the action no longer exists there.
					for _, artifact := range action.InputArtifacts {
						s.requiredInputs[artifact].Remove(OnDemandStage)
						s.requiredInputs[artifact].Add(firstOccurance)
					}
				}
			}
			if !matched {
				err := fmt.Errorf("one or more outputs from on-demand action '%s' could not be matched to an action in a fixed activity state. this action will be excluded from the action plan", action.Name)
				log.Warn("", "err", err)
				assignments[OnDemandStage].Remove(action)
			}
		}
	}
	return nil
}

func (s *scheduler) sortActions(assignments map[Stage]mapset.Set[*Action]) map[Stage][]*Action {
	sortedAssignments := map[Stage][]*Action{}
	assignedOutputs := []Artifact{}
	for stageIndex := range actionStages {
		stage := Stage(stageIndex)
		if actions, ok := assignments[stage]; ok {
			for {
				if actions.Cardinality() == 0 {
					break
				}
			ActionLoop:
				for _, action := range actions.ToSlice() {
					if len(action.InputArtifacts) == 0 {
						sortedAssignments[stage] = append(sortedAssignments[stage], action)
					} else {
						for _, input := range action.InputArtifacts {
							var assignedOutputMatchesInput bool
							for _, output := range assignedOutputs {
								if input == output {
									assignedOutputMatchesInput = true
								}
							}
							if !assignedOutputMatchesInput {
								continue ActionLoop
							}
						}
						sortedAssignments[stage] = append(sortedAssignments[stage], action)
					}
					assignedOutputs = append(assignedOutputs, action.OutputArtifacts...)
					actions.Remove(action)
				}
			}
		}
	}
	return sortedAssignments
}

func (s *scheduler) schedule(actions []string) (map[Stage][]*Action, error) {
	assignments := s.assignActivityStage(actions)
	s.bindActionInputs(assignments)
	if err := s.assignOnDemandActions(assignments); err != nil {
		return nil, err
	}
	return s.sortActions(assignments), nil
}

func (ap *ActionPlan) AddAction(spec ActionSpec, fields []string) {
	ap.Actions = append(ap.Actions, spec)
	ap.Fields = append(ap.Fields, fields...)
}

func (ap *ActionPlan) ToJson() (string, error) {
	data, err := json.Marshal(ap)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (ap *ActionPlan) runAction(source string, action *Action, client *dagger.Client) error {
	log.Info("Queuing action", "action", action.Name)
	container := client.Container().From(action.Image)
	container = container.WithMountedDirectory("/src", client.Host().Directory(source)).WithWorkdir("/src")
	for _, path := range action.Caches {
		container = container.WithMountedCache(path, client.CacheVolume(ap.id+path))
	}
	return action.Script(container, ap.Inputs, newActionUtilities(client, ap.artifacts))
}

func NewActionPlan(client *dagger.Client) *ActionPlan {
	plan := &ActionPlan{
		Inputs: make(map[string]interface{}),
		id:     time.Now().Format(time.RFC3339),
	}
	if client != nil {
		plan.artifacts = newArtifactStore(client)
	}
	return plan
}

func Run(source, spec string, client *dagger.Client, stages []string) error {
	ap := NewActionPlan(client)
	defer func() {
		for _, mount := range ap.artifacts.mounts {
			mount.Close()
		}
	}()
	if err := json.Unmarshal([]byte(spec), &ap); err != nil {
		return err
	}
	for _, input := range ap.Inputs {
		if input == nil {
			return errors.New("action plan cannot contain null user-defined inputs")
		}
	}
	actions := []string{}
	for _, a := range ap.Actions {
		actions = append(actions, a.Name)
	}
	schedule, err := newScheduler().schedule(actions)
	if err != nil {
		return err
	}
	stages = append([]string{"null"}, stages...)
	for stage, actions := range schedule {
		for _, stg := range stages {
			if stg == actionStages[stage] {
				for _, action := range actions {
					log.Info(fmt.Sprintf("> %s", action.Name))
				}
			}
		}
	}
	for stageId, key := range actionStages {
		for _, name := range stages {
			if key == name {
				stage := Stage(stageId)
				if actions, ok := schedule[stage]; ok {
					for _, action := range actions {
						if err := ap.runAction(source, action, client); err != nil {
							return err
						}
					}
				}
			}
		}
	}
	return nil
}

var registeredActions = map[string]*Action{}

func RegisterAction(action *Action) {
	registeredActions[action.Name] = action
}