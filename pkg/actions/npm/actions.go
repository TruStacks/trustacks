package npm

import (
	"context"
	"path/filepath"

	"dagger.io/dagger"
	"github.com/trustacks/trustacks/pkg/engine"
)

var npmTestAction = &engine.Action{
	Name:        "npmTest",
	DisplayName: "Npm Test",
	Description: "Run the test suite with npm test.",
	Image:       func(_ *engine.Config) string { return "node:alpine" },
	Stage:       engine.FeedbackStage,
	Caches:      []string{"/src/node_modules"},
	Script: func(container *dagger.Container, _ map[string]interface{}, _ *engine.ActionUtilities) error {
		container = container.WithExec([]string{"apk", "add", "bash"})
		container = container.WithExec([]string{"npm", "install"})
		container = container.WithEnvVariable("CI", "true")
		container = container.WithExec([]string{"npm", "test", "--coverage"})
		_, err := container.Stdout(context.Background())
		return err
	},
	AdmissionCriteria: []engine.Fact{NpmTestExistsFact},
}

var npmBuildAction = &engine.Action{
	Name:        "npmBuild",
	DisplayName: "Npm Build",
	Description: "Build the application with npm run build.",
	Image:       func(_ *engine.Config) string { return "node:alpine" },
	Stage:       engine.OnDemandStage,
	Caches:      []string{"/src/node_modules"},
	OutputArtifacts: []engine.Artifact{
		engine.ApplicationDistArtifact,
	},
	Script: func(container *dagger.Container, _ map[string]interface{}, utils *engine.ActionUtilities) error {
		container = container.WithExec([]string{"apk", "add", "bash"})
		container = container.WithExec([]string{"npm", "install"})
		container = container.WithExec([]string{"npm", "run", "build"})
		return utils.Export(container, engine.ApplicationDistArtifact, filepath.Join("/src", "build"))
	},
	AdmissionCriteria: []engine.Fact{NpmBuildExistsFact},
}

func init() {
	engine.RegisterPatternMatches([]engine.PatternMatch{
		{Kind: engine.FilePatternMatch, Pattern: ".*.test.js"},
		{Kind: engine.FilePatternMatch, Pattern: ".*.test.jsx"},
		{Kind: engine.FilePatternMatch, Pattern: ".*.test.ts"},
		{Kind: engine.FilePatternMatch, Pattern: ".*.test.tsx"},
	})
	engine.RegisterAction(npmTestAction)
	engine.RegisterAction(npmBuildAction)
}
