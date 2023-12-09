package engine

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"dagger.io/dagger"
	"github.com/briandowns/spinner"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
)

type ActionPlan struct {
	Actions   []string `json:"actions"`
	vars      map[string]interface{}
	id        string
	artifacts *ArtifactStore
}

func (ap *ActionPlan) AddAction(name string) {
	ap.Actions = append(ap.Actions, name)
}

func (ap *ActionPlan) ToJSON() (string, error) {
	data, err := json.Marshal(ap)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (ap *ActionPlan) prepare(spec string, client *dagger.Client) error {
	ap.artifacts = newArtifactStore(client)
	return json.Unmarshal([]byte(spec), &ap)
}

func (ap *ActionPlan) stageActions(stages []string) []string {
	actions := []string{}
	for _, stage := range stages {
		for _, name := range ap.Actions {
			action := registeredActions[name]
			if GetStage(action.Stage) == stage {
				actions = append(actions, action.Name)
			}
		}
	}
	return actions
}

func (ap *ActionPlan) checkInputs(actions []string) error {
	missingInputs := []string{}
	for _, name := range actions {
		for _, input := range registeredActions[name].Inputs {
			value := os.Getenv(string(input))
			if value == "" {
				missingInputs = append(missingInputs, string(input))
			} else {
				ap.vars[string(input)] = value
			}
		}
	}
	if len(missingInputs) > 0 {
		return fmt.Errorf(
			"the following inputs are required in order to run the action plan: %s.\nRun 'tsctl explain' to view inputs",
			strings.Join(missingInputs, ", "),
		)
	}
	return nil
}
func (ap *ActionPlan) logAction(action string) func(error) error {
	if os.Getenv("VERBOSE") == "true" {
		return func(err error) error { return err }
	}
	spin := spinner.New(spinner.CharSets[9], time.Duration(100)*time.Millisecond) //nolint:gomnd
	spin.Suffix = " " + action
	spin.Start()
	return func(err error) error {
		spin.Stop()
		status := lipgloss.NewStyle().Foreground(lipgloss.Color("#76CABC")).Render("✔")
		if err != nil {
			status = lipgloss.NewStyle().Foreground(lipgloss.Color("#df4053")).Render("✖")
		}
		fmt.Println(status + " " + action)
		return err
	}
}

func (ap *ActionPlan) runAction(source string, action *Action, client *dagger.Client, config *Config) error {
	stopLogger := ap.logAction(action.DisplayName)
	container := client.Pipeline(action.Name).Container().From(action.Image(config))
	container = container.WithMountedDirectory("/src", client.Host().Directory(source)).WithWorkdir("/src")
	for _, path := range action.Caches {
		container = container.WithMountedCache(path, client.CacheVolume(ap.id+path))
	}
	err := action.Script(container, ap.vars, newActionUtilities(client, ap.artifacts, config))
	return stopLogger(err)
}

func (ap *ActionPlan) close() {
	for _, mount := range ap.artifacts.mounts {
		mount.Close()
	}
}

func NewActionPlan() *ActionPlan {
	return &ActionPlan{
		vars: make(map[string]interface{}),
		id:   time.Now().Format(time.RFC3339),
	}
}

type RunArgs struct {
	Source              string
	Spec                string
	Client              *dagger.Client
	Stages              []string
	IgnoreMissingInputs bool
}

func Run(args RunArgs) error {
	ap := NewActionPlan()
	if err := ap.prepare(args.Spec, args.Client); err != nil {
		return err
	}
	defer ap.close()
	// stage "" is a placeholder for on-demand actions.
	stages := append([]string{""}, args.Stages...)
	actions := ap.stageActions(stages)
	if !args.IgnoreMissingInputs {
		if err := ap.checkInputs(actions); err != nil {
			return err
		}
	}
	schedule, err := newScheduler().schedule(actions)
	if err != nil {
		return err
	}
	for actionStage, actions := range schedule {
		for _, stageID := range stages {
			if stageID == actionStages[actionStage] {
				for _, action := range actions {
					if os.Getenv("DEBUG") != "" {
						log.Info(fmt.Sprintf("> %s", action.Name))
					}
				}
			}
		}
	}
	config, err := NewConfig()
	if err != nil {
		return err
	}
	for stageID, key := range actionStages {
		for _, name := range stages {
			if key == name {
				stage := Stage(stageID)
				if actions, ok := schedule[stage]; ok {
					for _, action := range actions {
						if err := ap.runAction(args.Source, action, args.Client, config); err != nil {
							return err
						}
					}
				}
			}
		}
	}
	return nil
}
