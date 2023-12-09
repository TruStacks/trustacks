package flake8

import (
	"context"
	"fmt"
	"strings"

	"dagger.io/dagger"
	"github.com/trustacks/trustacks/pkg/actions/python"
	"github.com/trustacks/trustacks/pkg/engine"
)

var flake8RunAction = &engine.Action{
	Name:        "flake8Run",
	DisplayName: "Flake8 Run",
	Description: "Run the flake8 linter",
	Image: func(config *engine.Config) string {
		if config.Python.Version != "" {
			return fmt.Sprintf("python:%s", config.Python.Version)
		}
		return "python"
	},
	Stage: engine.CommitStage,
	Script: func(container *dagger.Container, _ map[string]interface{}, utils *engine.ActionUtilities) error {
		config := utils.GetConfig()
		container = container.WithExec([]string{"apt", "update"})
		container = container.WithExec([]string{"apt", "install", "gcc", "-y"})
		container = container.WithExec([]string{"apt", "install", "-y", strings.Join(config.Python.Libraries, " ")})
		container, err := python.InstallPythonDependencies(container)
		if err != nil {
			return err
		}
		if config.Python.DevRequirements != "" {
			container = container.WithExec([]string{"pip", "install", "-r", config.Python.DevRequirements})
		}
		container = container.WithExec([]string{"pip", "install", "pytest"})
		container = container.WithExec([]string{"pytest"})
		_, err = container.Stdout(context.Background())
		return err
	},
	AdmissionCriteria: []engine.Fact{Flake8ConfigExistsFact},
}

func init() {
	engine.RegisterAction(flake8RunAction)
}
