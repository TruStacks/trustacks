package actions

import (
	"context"

	"dagger.io/dagger"
	"github.com/trustacks/trustacks/pkg/engine"
	"github.com/trustacks/trustacks/pkg/engine/rules"
	"github.com/trustacks/trustacks/pkg/plan"
)

var eslintRunAction = &plan.Action{
	Name:   "eslintRun",
	Image:  func(_ *plan.Config) string { return "node:alpine" },
	Stage:  plan.FeedbackStage,
	Caches: []string{"/src/node_modules"},
	Script: func(container *dagger.Container, _ map[string]interface{}, utils *plan.ActionUtilities) error {
		container = container.WithExec([]string{"apk", "add", "bash"})
		container = container.WithExec([]string{"npm", "install"})
		container = container.WithExec([]string{"npx", "-y", "eslint", "./"})
		_, err := container.Stdout(context.Background())
		return err
	},
}

func init() {
	engine.RegisterAdmissionResolver(
		plan.ActionSpec{
			Name:        eslintRunAction.Name,
			DisplayName: "ESLint Run",
			Description: "Lint the source with ESLint.",
		},
		[]engine.Fact{rules.ESLintConfigExistsFact},
		nil,
		nil,
	)
	plan.RegisterAction(eslintRunAction)
}
