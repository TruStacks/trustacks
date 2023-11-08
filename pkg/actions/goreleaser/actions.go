package goreleaser

import (
	"context"

	"dagger.io/dagger"
	"github.com/mitchellh/mapstructure"
	"github.com/trustacks/trustacks/pkg/engine"
	"github.com/trustacks/trustacks/pkg/plan"
)

var goreleaserRelease = &plan.Action{
	Name:   "goreleaserRelease",
	Image:  func(_ *plan.Config) string { return "golang" },
	Stage:  plan.ReleaseStage,
	Caches: []string{"/go/pkg/mod"},
	Script: func(container *dagger.Container, inputs map[string]interface{}, _ *plan.ActionUtilities) error {
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
}

func init() {
	engine.RegisterAdmissionResolver(
		plan.ActionSpec{
			Name:        goreleaserRelease.Name,
			DisplayName: "Goreleaser Release",
			Description: "Release the golang application with goreleaser.",
		},
		[]engine.Fact{GoreleaserConfigExistsFact},
		nil,
		[]string{string(plan.GithubToken)},
	)
	plan.RegisterAction(goreleaserRelease)
}
