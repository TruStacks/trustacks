package eslint

import (
	"context"

	"dagger.io/dagger"
	"github.com/trustacks/trustacks/pkg/engine"
)

var eslintRunAction = &engine.Action{
	Name:        "eslintRun",
	DisplayName: "ESLint Run",
	Description: "Lint the source with ESLint.",
	Image:       func(_ *engine.Config) string { return "node:alpine" },
	Stage:       engine.CommitStage,
	Caches:      []string{"/src/node_modules"},
	Script: func(container *dagger.Container, _ map[string]interface{}, utils *engine.ActionUtilities) error {
		container = container.WithExec([]string{"apk", "add", "bash"})
		container = container.WithExec([]string{"npm", "install"})
		container = container.WithExec([]string{"npx", "-y", "eslint", "./"})
		_, err := container.Sync(context.Background())
		return err
	},
	AdmissionCriteria: []engine.Fact{ESLintConfigExistsFact},
}

func init() {
	engine.RegisterAction(eslintRunAction)
}
