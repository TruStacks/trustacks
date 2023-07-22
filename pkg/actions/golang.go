package actions

import (
	"context"

	"dagger.io/dagger"
	"github.com/trustacks/pkg/engine"
	"github.com/trustacks/pkg/plan"
)

var golangTest = &plan.Action{
	Name:   "golangTest",
	Image:  "golang",
	State:  plan.FeebackState,
	Caches: []string{"/go/pkg/mod"},
	Script: func(container *dagger.Container, _ map[string]interface{}, _ *plan.ActionUtilities) error {
		container = container.WithExec([]string{"go", "test", "./...", "-v", "-short", "-cover"})
		_, err := container.Stdout(context.Background())
		return err
	},
}

func init() {
	engine.RegisterPatternMatches([]engine.PatternMatch{
		{
			Kind:       engine.FilePatternMatch,
			Pattern:    "_test.go",
			Exclusions: &[]string{"testdata", "vendor"},
		},
	})
	engine.RegisterAdmissionResolver(golangTest.Name, []engine.Fact{engine.HasGolangTestSourcesFact}, nil)
	plan.RegisterAction(golangTest)
}
