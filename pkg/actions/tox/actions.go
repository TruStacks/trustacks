package tox

import (
	"context"

	"dagger.io/dagger"
	"github.com/trustacks/trustacks/pkg/actions/python"
	"github.com/trustacks/trustacks/pkg/engine"
)

var toxRunAction = &engine.Action{
	Name:        "toxRun",
	DisplayName: "Tox Run",
	Description: "Run the python test suite using tox",
	Image:       func(_ *engine.Config) string { return "python" },
	Stage:       engine.CommitStage,
	Script: func(container *dagger.Container, _ map[string]interface{}, _ *engine.ActionUtilities) error {
		container, err := python.InstallPythonDependencies(container)
		if err != nil {
			return err
		}
		container = container.WithExec([]string{"pip", "install", "tox"})
		container = container.WithExec([]string{"tox", "run"})
		_, err = container.Sync(context.Background())
		return err

	},
	AdmissionCriteria: []engine.Fact{ToxIniExistsFact},
}

func init() {
	engine.RegisterPatternMatches([]engine.PatternMatch{
		{Kind: engine.FilePatternMatch, Pattern: "test_.*.py"},
	})
	engine.RegisterAction(toxRunAction)
}
