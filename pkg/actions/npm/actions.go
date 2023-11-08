package npm

import (
	"context"
	"path/filepath"

	"dagger.io/dagger"
	"github.com/trustacks/trustacks/pkg/engine"
	"github.com/trustacks/trustacks/pkg/plan"
)

var npmTestAction = &plan.Action{
	Name:   "npmTest",
	Image:  func(_ *plan.Config) string { return "node:alpine" },
	Stage:  plan.FeedbackStage,
	Caches: []string{"/src/node_modules"},
	Script: func(container *dagger.Container, _ map[string]interface{}, _ *plan.ActionUtilities) error {
		container = container.WithExec([]string{"apk", "add", "bash"})
		container = container.WithExec([]string{"npm", "install"})
		container = container.WithEnvVariable("CI", "true")
		container = container.WithExec([]string{"npm", "test", "--coverage"})
		_, err := container.Stdout(context.Background())
		return err
	},
}

var npmBuildAction = &plan.Action{
	Name:   "npmBuild",
	Image:  func(_ *plan.Config) string { return "node:alpine" },
	Stage:  plan.OnDemandStage,
	Caches: []string{"/src/node_modules"},
	OutputArtifacts: []plan.Artifact{
		plan.ApplicationDistArtifact,
	},
	Script: func(container *dagger.Container, _ map[string]interface{}, utils *plan.ActionUtilities) error {
		container = container.WithExec([]string{"apk", "add", "bash"})
		container = container.WithExec([]string{"npm", "install"})
		container = container.WithExec([]string{"npm", "run", "build"})
		return utils.Export(container, plan.ApplicationDistArtifact, filepath.Join("/src", "build"))
	},
}

func init() {
	engine.RegisterPatternMatches([]engine.PatternMatch{
		{Kind: engine.FilePatternMatch, Pattern: ".*.test.js"},
		{Kind: engine.FilePatternMatch, Pattern: ".*.test.jsx"},
		{Kind: engine.FilePatternMatch, Pattern: ".*.test.ts"},
		{Kind: engine.FilePatternMatch, Pattern: ".*.test.tsx"},
	})
	engine.RegisterAdmissionResolver(
		plan.ActionSpec{
			Name:        npmTestAction.Name,
			DisplayName: "Npm Test",
			Description: "Run the test suite with npm test.",
		},
		[]engine.Fact{NpmTestExistsFact},
		nil,
		nil,
	)
	plan.RegisterAction(npmTestAction)

	engine.RegisterAdmissionResolver(
		plan.ActionSpec{
			Name:        npmBuildAction.Name,
			DisplayName: "Npm Build",
			Description: "Build the application with npm run build.",
		},
		[]engine.Fact{NpmBuildExistsFact},
		nil,
		nil,
	)
	plan.RegisterAction(npmBuildAction)
}
