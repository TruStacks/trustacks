package tox

import (
	"context"

	"dagger.io/dagger"
	"github.com/trustacks/trustacks/pkg/actions/python"
	"github.com/trustacks/trustacks/pkg/engine"
	"github.com/trustacks/trustacks/pkg/plan"
)

var toxRunAction = &plan.Action{
	Name:  "toxRun",
	Image: func(_ *plan.Config) string { return "python" },
	Stage: plan.FeedbackStage,
	Script: func(container *dagger.Container, _ map[string]interface{}, _ *plan.ActionUtilities) error {
		container, err := python.InstallPythonDependencies(container)
		if err != nil {
			return err
		}
		container = container.WithExec([]string{"pip", "install", "tox"})
		container = container.WithExec([]string{"tox", "run"})
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
			Name:        toxRunAction.Name,
			DisplayName: "Tox Run",
			Description: "Run the python test suite using tox",
		},
		[]engine.Fact{ToxIniExistsFact},
		nil,
		nil,
	)
	plan.RegisterAction(toxRunAction)
}
