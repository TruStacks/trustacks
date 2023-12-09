package actions

import (
	// import actions
	_ "github.com/trustacks/trustacks/pkg/actions/argocd"
	_ "github.com/trustacks/trustacks/pkg/actions/container"
	_ "github.com/trustacks/trustacks/pkg/actions/eslint"
	_ "github.com/trustacks/trustacks/pkg/actions/golang"
	_ "github.com/trustacks/trustacks/pkg/actions/golangcilint"
	_ "github.com/trustacks/trustacks/pkg/actions/goreleaser"
	_ "github.com/trustacks/trustacks/pkg/actions/javascript"
	_ "github.com/trustacks/trustacks/pkg/actions/npm"
	_ "github.com/trustacks/trustacks/pkg/actions/pytest"
	_ "github.com/trustacks/trustacks/pkg/actions/python"
	_ "github.com/trustacks/trustacks/pkg/actions/sonarqube"
	_ "github.com/trustacks/trustacks/pkg/actions/tox"
	_ "github.com/trustacks/trustacks/pkg/actions/trivy"
)
