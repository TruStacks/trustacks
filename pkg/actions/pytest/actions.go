package pytest

import (
	"context"
	"fmt"
	"strings"

	"dagger.io/dagger"
	"github.com/trustacks/trustacks/pkg/actions/python"
	"github.com/trustacks/trustacks/pkg/actions/tox"
	"github.com/trustacks/trustacks/pkg/engine"
	"github.com/trustacks/trustacks/pkg/plan"
)

var pytestRunAction = &plan.Action{
	Name: "pytestRun",
	Image: func(config *plan.Config) string {
		if config.Python.Version != "" {
			return fmt.Sprintf("python:%s", config.Python.Version)
		}
		return "python"
	},
	Stage: plan.FeedbackStage,
	Script: func(container *dagger.Container, _ map[string]interface{}, utils *plan.ActionUtilities) error {
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
}

func init() {
	engine.RegisterPatternMatches([]engine.PatternMatch{
		{Kind: engine.FilePatternMatch, Pattern: "test_.*.py"},
	})
	engine.RegisterAdmissionResolver(
		plan.ActionSpec{
			Name:        pytestRunAction.Name,
			DisplayName: "PyTest Run",
			Description: "Run the python test suite using pytest",
		},
		[]engine.Fact{PytestDependencyExistsFact},
		[]engine.Fact{tox.ToxIniExistsFact},
		nil,
	)
	plan.RegisterAction(pytestRunAction)
}
