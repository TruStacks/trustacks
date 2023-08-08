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
	Stage:  plan.OnDemandStage,
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
	engine.RegisterAdmissionResolver(
		plan.ActionSpec{
			Name:        packageJsonVersion.Name,
			DisplayName: "Package JSON Version",
			Description: "Use the package.json version as the semantic release version for versioned application artifacts.",
		},
		[]engine.Fact{engine.PackageJsonVersionExistsFact},
		nil,
	)
	plan.RegisterAction(packageJsonVersion)
}
