package actions

import (
	"context"

	"dagger.io/dagger"
	"trustacks.io/trustacks/plan"
)

var golangTest = &plan.Action{
	Name:   "golangTest",
	Image:  "golang",
	Stage:  plan.FeedbackStage,
	Caches: []string{"/go/pkg/mod"},
	Script: func(container *dagger.Container, _ map[string]interface{}, _ *plan.ActionUtilities) error {
		container = container.WithExec([]string{"go", "test", "./...", "-v", "-short", "-cover"})
		_, err := container.Stdout(context.Background())
		return err
	},
}

// func init() {
// 	engine.RegisterPatternMatches([]engine.PatternMatch{
// 		{
// 			Kind:       engine.FilePatternMatch,
// 			Pattern:    "_test.go",
// 			Exclusions: &[]string{"testdata", "vendor"},
// 		},
// 	})
// 	engine.RegisterAdmissionResolver(
// 		plan.ActionSpec{
// 			Name:        golangTest.Name,
// 			DisplayName: "Golang Test",
// 			Description: "Run the test suite with go test.",
// 		},
// 		[]engine.Fact{},
// 		nil,
// 	)
// 	plan.RegisterAction(golangTest)
// }
