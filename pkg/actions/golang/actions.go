package golang

import (
	"context"

	"dagger.io/dagger"
	"github.com/trustacks/trustacks/pkg/engine"
)

var golangTest = &engine.Action{
	Name:        "golangTest",
	DisplayName: "Golang Test",
	Description: "Run the test suite with go test.",
	Image:       func(_ *engine.Config) string { return "golang" },
	Stage:       engine.FeedbackStage,
	Caches:      []string{"/go/pkg/mod"},
	Script: func(container *dagger.Container, _ map[string]interface{}, _ *engine.ActionUtilities) error {
		container = container.WithExec([]string{"go", "test", "./...", "-v", "-short", "-cover"})
		_, err := container.Stdout(context.Background())
		return err
	},
	AdmissionCriteria: []engine.Fact{GolangTestsExistFact},
}

func init() {
	engine.RegisterPatternMatches([]engine.PatternMatch{
		{
			Kind:       engine.FilePatternMatch,
			Pattern:    ".*_test.go",
			Exclusions: &[]string{"testdata", "vendor"},
		},
	})
	engine.RegisterAction(golangTest)
}
