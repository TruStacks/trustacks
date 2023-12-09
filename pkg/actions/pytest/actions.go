package pytest

import (
	"context"
	"fmt"
	"strings"

	"dagger.io/dagger"
	"github.com/trustacks/trustacks/pkg/actions/python"
	"github.com/trustacks/trustacks/pkg/actions/tox"
	"github.com/trustacks/trustacks/pkg/engine"
)

var pytestRunAction = &engine.Action{
	Name:        "pytestRun",
	DisplayName: "PyTest Run",
	Description: "Run the python test suite using pytest",
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
		_, err = container.Sync(context.Background())
		return err
	},
	AdmissionCriteria: []engine.Fact{PytestDependencyExistsFact},
	ExclusionCriteria: []engine.Fact{tox.ToxIniExistsFact},
}

func init() {
	engine.RegisterPatternMatches([]engine.PatternMatch{
		{Kind: engine.FilePatternMatch, Pattern: "test_.*.py"},
	})
	engine.RegisterAction(pytestRunAction)
}
