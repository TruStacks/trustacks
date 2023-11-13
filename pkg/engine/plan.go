package engine

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"dagger.io/dagger"
	"github.com/charmbracelet/log"
)

type ActionPlan struct {
	Actions   []string     `json:"actions"`
	Inputs    []InputField `json:"inputs,omitempty"`
	vars      map[string]interface{}
	id        string
	artifacts *ArtifactStore
}

func NewActionPlan(client *dagger.Client) *ActionPlan {
	plan := &ActionPlan{
		vars: make(map[string]interface{}),
		id:   time.Now().Format(time.RFC3339),
	}
	if client != nil {
		plan.artifacts = newArtifactStore(client)
	}
	return plan
}

func (ap *ActionPlan) AddAction(name string, inputs []InputField) {
	ap.Actions = append(ap.Actions, name)
	ap.Inputs = append(ap.Inputs, inputs...)
}

func (ap *ActionPlan) ToJson() (string, error) {
	data, err := json.Marshal(ap)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (ap *ActionPlan) runAction(source string, action *Action, client *dagger.Client, config *Config) error {
	log.Info("Queuing action", "action", action.Name)
	container := client.Pipeline(action.Name).Container().From(action.Image(config))
	container = container.WithMountedDirectory("/src", client.Host().Directory(source)).WithWorkdir("/src")
	for _, path := range action.Caches {
		container = container.WithMountedCache(path, client.CacheVolume(ap.id+path))
	}
	return action.Script(container, ap.vars, newActionUtilities(client, ap.artifacts, config))
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
	missingInputs := []string{}
	for _, input := range ap.Inputs {
		stagedActionInputs := []string{}
		for _, name := range ap.Actions {
			for _, stage := range stages {
				fmt.Println(name)
				action, ok := registeredActions[name]
				if ok && GetStage(action.Stage) == stage {
					fmt.Println(action.Inputs)
				}
			}
		}
		fmt.Println(stagedActionInputs)
		value := os.Getenv(string(input))
		if value == "" {
			missingInputs = append(missingInputs, string(input))
		} else {
			ap.vars[string(input)] = value
		}
	}
	if len(missingInputs) > 0 {
		return fmt.Errorf(
			"the following inpust are required in order to run the action plan: %s.\nRun 'tsctl explain' to view inputs",
			strings.Join(missingInputs, ", "),
		)
	}
	actions := ap.Actions
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
	config, err := NewConfig()
	if err != nil {
		return err
	}
	for stageId, key := range actionStages {
		for _, name := range stages {
			if key == name {
				stage := Stage(stageId)
				if actions, ok := schedule[stage]; ok {
					for _, action := range actions {
						if err := ap.runAction(source, action, client, config); err != nil {
							return err
						}
					}
				}
			}
		}
	}
	return nil
}
