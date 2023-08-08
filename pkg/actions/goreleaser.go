package actions

import (
	"context"

	"dagger.io/dagger"
	"github.com/trustacks/pkg/engine"
	"github.com/trustacks/pkg/plan"
)

var goreleaserPreRelease = &plan.Action{
	Name:   "goreleaserBuild",
	Image:  "golang:alpine",
	Stage:  plan.StageStage,
	Caches: []string{"/go/pkg/mod"},
	Script: func(container *dagger.Container, _ map[string]interface{}, _ *plan.ActionUtilities) error {
		container = container.WithExec([]string{"apk", "add", "curl", "git", "docker-cli"})
		container = container.WithExec([]string{"curl", "-L", "-O", "https://github.com/goreleaser/goreleaser/releases/download/v1.18.2/goreleaser_1.18.2_x86.apk"})
		container = container.WithExec([]string{"apk", "add", "goreleaser_1.18.2_x86.apk", "--allow-untrusted"})
		container = container.WithExec([]string{"goreleaser", "build", "--snapshot", "--clean"})
		_, err := container.Stdout(context.Background())
		return err
	},
}

func init() {
	engine.RegisterAdmissionResolver(
		plan.ActionSpec{
			Name:        goreleaserPreRelease.Name,
			DisplayName: "GoReleaser Pre-Release",
			Description: "Run goreleaser with the prerelease flag.",
		},
		[]engine.Fact{engine.GoreleaserExistsFact},
		nil,
	)
	plan.RegisterAction(goreleaserPreRelease)
}
