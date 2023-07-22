package actions

import (
	"context"
	"path/filepath"

	"dagger.io/dagger"
	"github.com/trustacks/pkg/engine"
	"github.com/trustacks/pkg/plan"
)

// reactScriptsTestAction .
var reactScriptsTestAction = &plan.Action{
	Name:   "reactScriptsTest",
	Image:  "node:alpine",
	State:  plan.FeebackState,
	Caches: []string{"/src/node_modules"},
	Script: func(container *dagger.Container, _ map[string]interface{}, _ *plan.ActionUtilities) error {
		container = container.WithExec([]string{"apk", "add", "bash"})
		container = container.WithExec([]string{"npm", "install"})
		container = container.WithEnvVariable("CI", "true")
		container = container.WithExec([]string{"npx", "react-scripts", "test", "--coverage"})
		_, err := container.Stdout(context.Background())
		return err
	},
}

// reactScriptsBuildAction .
var reactScriptsBuildAction = &plan.Action{
	Name:   "reactScriptsBuild",
	Image:  "node:alpine",
	State:  plan.OnDemandState,
	Caches: []string{"/src/node_modules"},
	OutputArtifacts: []plan.Artifact{
		plan.ApplicationDistArtifact,
	},
	Script: func(container *dagger.Container, _ map[string]interface{}, utils *plan.ActionUtilities) error {
		container = container.WithExec([]string{"apk", "add", "bash"})
		container = container.WithExec([]string{"npm", "install"})
		container = container.WithExec([]string{"npx", "react-scripts", "build"})
		return utils.Export(container, plan.ApplicationDistArtifact, filepath.Join("/src", "build"))
	},
}

func init() {
	engine.RegisterAdmissionResolver(reactScriptsTestAction.Name, []engine.Fact{engine.ReactScriptsTestExistsFact}, nil)
	plan.RegisterAction(reactScriptsTestAction)

	engine.RegisterAdmissionResolver(reactScriptsBuildAction.Name, []engine.Fact{engine.ReactScriptsBuildExistsFact}, nil)
	plan.RegisterAction(reactScriptsBuildAction)
}
