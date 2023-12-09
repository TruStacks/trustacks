package golang

import (
	"context"
	"fmt"

	"dagger.io/dagger"
	"github.com/trustacks/trustacks/pkg/engine"
)

const imageName = "golang"

var golangBuild = &engine.Action{
	Name:        "golangBuild",
	DisplayName: "Golang Build",
	Description: "Build the golang application.",
	Image:       func(config *engine.Config) string { return imageName },
	Stage:       engine.CommitStage,
	Caches:      []string{"/go/pkg/mod"},
	OutputArtifacts: []engine.Artifact{
		engine.BuildArtifact,
	},
	Script: func(container *dagger.Container, _ map[string]interface{}, utils *engine.ActionUtilities) error {
		entries, err := container.Directory("./cmd").Entries(context.Background())
		if err != nil {
			return err
		}
		version := utils.GetConfig().Common.Version
		if version != "" {
			container = container.WithEnvVariable("VERSION", version)
		}
		for _, entry := range entries {
			container = container.WithExec([]string{
				"go",
				"build",
				"-ldflags",
				utils.GetConfig().Golang.LDFlags,
				"-o",
				fmt.Sprintf(".build/%s", entry),
				fmt.Sprintf("./cmd/%s", entry),
			})
		}
		if err := utils.Export(container, engine.BuildArtifact, ".build"); err != nil {
			return err
		}
		_, err = container.Sync(context.Background())
		return err
	},
	AdmissionCriteria: []engine.Fact{
		GoModExistsFact,
		GolangCmdExistsFact,
	},
}

var golangTest = &engine.Action{
	Name:        "golangTest",
	DisplayName: "Golang Test",
	Description: "Run the unit test suite with go test.",
	Image:       func(config *engine.Config) string { return imageName },
	Stage:       engine.CommitStage,
	Caches:      []string{"/go/pkg/mod"},
	OutputArtifacts: []engine.Artifact{
		engine.CoverageArtifact,
	},
	Script: func(container *dagger.Container, _ map[string]interface{}, utils *engine.ActionUtilities) error {
		container = container.WithExec([]string{"go", "test", "./...", "-v", "-short", "-coverprofile", "coverage.out"})
		if err := utils.Export(container, engine.CoverageArtifact, "coverage.out"); err != nil {
			return err
		}
		_, err := container.Sync(context.Background())
		return err
	},
	AdmissionCriteria: []engine.Fact{GolangTestsExistsFact},
}

var golangIntegrationTest = &engine.Action{
	Name:        "golangIntegrationTest",
	DisplayName: "Golang Integration Test",
	Description: "Run the integration test suite with go test",
	Image:       func(_ *engine.Config) string { return imageName },
	Stage:       engine.CommitStage,
	Caches:      []string{"/go/pkg/mod"},
	Script: func(container *dagger.Container, _ map[string]interface{}, utils *engine.ActionUtilities) error {
		container, stop, err := utils.WithDockerdService(container)
		if err != nil {
			return err
		}
		defer stop()
		container = utils.WithDockerCLI(engine.DockerCLIOnDebian, container)
		container = container.WithExec([]string{"go", "test", "./...", "-v", "-run", "Integration"})
		_, err = container.Sync(context.Background())
		return err
	},
	AdmissionCriteria: []engine.Fact{GolangIntegrationTestsExistsFact},
}

func init() {
	engine.RegisterPatternMatches([]engine.PatternMatch{
		{
			Kind:       engine.FilePatternMatch,
			Pattern:    ".*_test.go",
			Exclusions: &[]string{"testdata", "vendor"},
		},
	})
	engine.RegisterAction(golangBuild)
	engine.RegisterAction(golangTest)
	engine.RegisterAction(golangIntegrationTest)
}
