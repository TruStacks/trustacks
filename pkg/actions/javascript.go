package actions

import (
	"path/filepath"

	"dagger.io/dagger"
	"github.com/trustacks/pkg/engine"
	"github.com/trustacks/pkg/plan"
)

var packageJsonVersion = &plan.Action{
	Name:   "packageJsonVersion",
	Image:  "node:alpine",
	State:  plan.OnDemandState,
	Caches: []string{"/src/node_modules"},
	OutputArtifacts: []plan.Artifact{
		plan.SemanticVersionArtifact,
	},
	Script: func(container *dagger.Container, _ map[string]interface{}, utils *plan.ActionUtilities) error {
		container = container.WithExec([]string{"apk", "add", "jq"})
		container = container.WithExec([]string{"/bin/sh", "-c", "cat package.json | jq '.version' -r > /tmp/version"})
		return utils.Export(container, plan.SemanticVersionArtifact, filepath.Join("/tmp", "version"))
	},
}

func init() {
	engine.RegisterAdmissionResolver(packageJsonVersion.Name, []engine.Fact{engine.PackageJsonVersionExistsFact}, nil)
	plan.RegisterAction(packageJsonVersion)
}
