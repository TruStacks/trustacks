package actions

import (
	"context"
	"path/filepath"

	"dagger.io/dagger"
	"github.com/trustacks/trustacks/pkg/engine"
	"github.com/trustacks/trustacks/pkg/engine/rules"
	"github.com/trustacks/trustacks/pkg/plan"
)

var reactScriptsTestAction = &plan.Action{
	Name:   "reactScriptsTest",
	Image:  "node:alpine",
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

var reactScriptsBuildAction = &plan.Action{
	Name:   "reactScriptsBuild",
	Image:  "node:alpine",
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
		{Kind: engine.FilePatternMatch, Pattern: `\.test.js`},
		{Kind: engine.FilePatternMatch, Pattern: `\.test.jsx`},
		{Kind: engine.FilePatternMatch, Pattern: `\.test.ts`},
		{Kind: engine.FilePatternMatch, Pattern: `\.test.tsx`},
	})
	engine.RegisterAdmissionResolver(
		plan.ActionSpec{
			Name:        reactScriptsTestAction.Name,
			DisplayName: "React Scripts Test",
			Description: "Run the test suite with react-scripts test.",
		},
		[]engine.Fact{rules.ReactScriptsTestExistsFact},
		nil,
	)
	plan.RegisterAction(reactScriptsTestAction)

	engine.RegisterAdmissionResolver(
		plan.ActionSpec{
			Name:        reactScriptsBuildAction.Name,
			DisplayName: "React Scripts Build",
			Description: "Build production react assets with react-scripts build.",
		},
		[]engine.Fact{rules.ReactScriptsBuildExistsFact},
		nil,
	)
	plan.RegisterAction(reactScriptsBuildAction)
}
