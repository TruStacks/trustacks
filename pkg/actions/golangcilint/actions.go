package golangcilint

import (
	"context"
	"fmt"

	"dagger.io/dagger"
	"github.com/trustacks/trustacks/pkg/engine"
)

const golangciLintVersion = "v1.53.3"

var golangCILintRun = &engine.Action{
	Name:        "golangCILintRun",
	DisplayName: "GolangCILint Run",
	Description: "Lint the source with golangci-lint.",
	Image:       func(_ *engine.Config) string { return "golang:alpine" },
	Stage:       engine.FeedbackStage,
	Caches:      []string{"/go/pkg/mod"},
	Script: func(container *dagger.Container, _ map[string]interface{}, _ *engine.ActionUtilities) error {
		container = container.WithExec([]string{"apk", "add", "bash", "curl", "git"})
		container = container.WithExec([]string{
			"/bin/sh",
			"-c",
			fmt.Sprintf("curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin %s", golangciLintVersion),
		})
		container = container.WithExec([]string{"golangci-lint", "run"})
		_, err := container.Stdout(context.Background())
		return err
	},
	AdmissionCriteria: []engine.Fact{GolangCILintConfigExistsFact},
}

func init() {
	engine.RegisterAction(golangCILintRun)
}
