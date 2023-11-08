package run

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"dagger.io/dagger"
	"github.com/charmbracelet/log"
	"github.com/trustacks/trustacks/pkg/engine"
	"github.com/trustacks/trustacks/pkg/plan"
)

type RunCmdOptions struct {
	Source     string
	Plan       string
	Stages     []string
	Prerelease bool
}

func removeReleaseStage(stages []string) []string {
	stageIndex := -1
	for i, stage := range stages {
		if stage == plan.GetStage(plan.ReleaseStage) {
			stageIndex = i
		}
	}
	if stageIndex >= 0 {
		stages = append(stages[:stageIndex], stages[stageIndex+1:]...)
	}
	return stages
}

func RunCmd(options *RunCmdOptions) error {
	var planData map[string]interface{}
	if _, err := os.Stat(options.Plan); os.IsNotExist(err) {
		spec, err := engine.New().CreateActionPlan(options.Source, false)
		if err != nil {
			return fmt.Errorf("failed creating the action plan: %s", err)
		}
		if err := json.Unmarshal([]byte(spec), &planData); err != nil {
			return fmt.Errorf("failed unmarshaling the action plan data: %s", err)
		}
	} else {
		planJson, err := os.ReadFile(options.Plan)
		if err != nil {
			return fmt.Errorf("failed opening plan file: %s", err)
		}
		if err := json.Unmarshal(planJson, &planData); err != nil {
			return fmt.Errorf("failed parsing plan file: %s", err)
		}
	}
	spec, err := json.Marshal(planData)
	if err != nil {
		return fmt.Errorf("failed converting plan file to spec: %s", err)
	}
	client, err := dagger.Connect(context.Background(), dagger.WithLogOutput(os.Stdout))
	if err != nil {
		return fmt.Errorf("failed connecting to the dagger agent")
	}
	if options.Prerelease {
		removeReleaseStage(options.Stages)
	}
	if err := plan.Run(options.Source, string(spec), client, options.Stages); err != nil {
		log.Error("", "err", err)
		os.Exit(1)
	}
	return nil
}
