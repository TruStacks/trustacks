package actions

import (
	"context"
	"fmt"

	"dagger.io/dagger"
	"github.com/trustacks/pkg/engine"
	"github.com/trustacks/pkg/plan"
)

var golangCILintRun = &plan.Action{
	Name:   "golangCILintRun",
	Image:  "golang:alpine",
	State:  plan.FeebackState,
	Caches: []string{"/go/pkg/mod"},
	Script: func(container *dagger.Container, _ map[string]interface{}, _ *plan.ActionUtilities) error {
		golangciLintVersion := "v1.53.3"
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
}

func init() {
	engine.RegisterAdmissionResolver(golangCILintRun.Name, []engine.Fact{engine.GolangCILintExistsFact}, nil)
	plan.RegisterAction(golangCILintRun)
}
