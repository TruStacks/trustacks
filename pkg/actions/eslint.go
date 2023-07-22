package actions

import (
	"context"

	"dagger.io/dagger"
	"github.com/trustacks/pkg/engine"
	"github.com/trustacks/pkg/plan"
)

var eslintRunAction = &plan.Action{
	Name:   "eslintRun",
	Image:  "node:alpine",
	State:  plan.FeebackState,
	Caches: []string{"/src/node_modules"},
	Script: func(container *dagger.Container, _ map[string]interface{}, _ *plan.ActionUtilities) error {
		container = container.WithExec([]string{"apk", "add", "bash"})
		container = container.WithExec([]string{"npm", "install"})
		container = container.WithExec([]string{"npx", "-y", "eslint", "./"})
		_, err := container.Stdout(context.Background())
		return err
	},
}

func init() {
	engine.RegisterAdmissionResolver(eslintRunAction.Name, []engine.Fact{engine.ReactScriptsTestExistsFact}, nil)
	plan.RegisterAction(eslintRunAction)
}
