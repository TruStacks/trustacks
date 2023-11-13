package goreleaser

import (
	"context"

	"dagger.io/dagger"
	"github.com/mitchellh/mapstructure"
	"github.com/trustacks/trustacks/pkg/engine"
)

var goreleaserRelease = &engine.Action{
	Name:        "goreleaserRelease",
	DisplayName: "Goreleaser Release",
	Description: "Release the golang application with goreleaser.",
	Image:       func(_ *engine.Config) string { return "golang" },
	Stage:       engine.ReleaseStage,
	Caches:      []string{"/go/pkg/mod"},
	Script: func(container *dagger.Container, inputs map[string]interface{}, _ *engine.ActionUtilities) error {
		args := struct {
			GITHUB_TOKEN string
		}{}
		if err := mapstructure.Decode(inputs, &args); err != nil {
			return err
		}
		container = container.WithExec([]string{"/bin/sh", "-c", "echo 'deb [trusted=yes] https://repo.goreleaser.com/apt/ /' | sudo tee /etc/apt/sources.list.d/goreleaser.list"})
		container = container.WithExec([]string{"apt", "update"})
		container = container.WithExec([]string{"apt", "install", "git", "goreleaser"})
		_, err := container.Stdout(context.Background())
		return err
	},
	Inputs: []engine.InputField{
		engine.GithubToken,
	},
	AdmissionCriteria: []engine.Fact{GoreleaserConfigExistsFact},
}

func init() {
	engine.RegisterAction(goreleaserRelease)
}
