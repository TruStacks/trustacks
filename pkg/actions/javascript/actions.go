package javascript

import (
	"path/filepath"

	"dagger.io/dagger"
	"github.com/trustacks/trustacks/pkg/engine"
)

var packageJSONVersion = &engine.Action{
	Name:        "packageJSONVersion",
	DisplayName: "Package JSON Version",
	Description: "Use the package.json version as the semantic release version for versioned application artifacts.",
	Image:       func(_ *engine.Config) string { return "node:alpine" },
	Stage:       engine.OnDemand,
	Caches:      []string{"/src/node_modules"},
	OutputArtifacts: []engine.Artifact{
		engine.SemanticVersionArtifact,
	},
	Script: func(container *dagger.Container, _ map[string]interface{}, utils *engine.ActionUtilities) error {
		container = container.WithExec([]string{"apk", "add", "jq"})
		container = container.WithExec([]string{"/bin/sh", "-c", "cat package.json | jq '.version' -r > /tmp/version"})
		return utils.Export(container, engine.SemanticVersionArtifact, filepath.Join("/tmp", "version"))
	},
	AdmissionCriteria: []engine.Fact{PackageJSONVersionExistsFact},
}

func init() {
	engine.RegisterAction(packageJSONVersion)
}
