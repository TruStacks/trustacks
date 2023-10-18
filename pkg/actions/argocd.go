package actions

import (
	"context"
	"strings"

	"dagger.io/dagger"
	"github.com/mitchellh/mapstructure"
	"github.com/trustacks/trustacks/pkg/engine"
	"github.com/trustacks/trustacks/pkg/engine/rules"
	"github.com/trustacks/trustacks/pkg/plan"
	"gopkg.in/yaml.v2"
)

var argocdSync = &plan.Action{
	Name:  "argocdSync",
	Image: func(_ *plan.Config) string { return "argoproj/argocd" },
	Stage: plan.StageStage,
	Script: func(container *dagger.Container, inputs map[string]interface{}, utils *plan.ActionUtilities) error {
		var err error
		args := struct {
			ARGOCD_SERVER     string
			ARGOCD_AUTH_TOKEN string
		}{}
		if err := mapstructure.Decode(inputs, &args); err != nil {
			return err
		}
		config := utils.GetConfig()
		container = container.WithSecretVariable("ARGOCD_AUTH_TOKEN", utils.SetSecret("argocdAuthToken", args.ARGOCD_AUTH_TOKEN))
		container = container.WithEnvVariable("ARGOCD_SERVER", args.ARGOCD_SERVER)
		container = container.WithExec([]string{"grep", "-r", "argoproj.io/v1alpha1"})
		stdout, err := container.Stdout(context.Background())
		if err != nil {
			return err
		}
		appSpecPath := strings.Split(strings.Split(stdout, "\n")[0], ":")[0]
		createCmd := []string{"argocd", "app", "create", "-f", appSpecPath, "--upsert"}
		if config.ArgoCD.Insecure {
			createCmd = append(createCmd, "--insecure")
		}
		if config.ArgoCD.GRPCWeb {
			createCmd = append(createCmd, "--grpc-web")
		}
		container = container.WithExec(createCmd)
		container = container.WithExec([]string{"cat", appSpecPath})
		stdout, err = container.Stdout(context.Background())
		if err != nil {
			return err
		}
		spec := map[string]interface{}{}
		if err := yaml.Unmarshal([]byte(stdout), &spec); err != nil {
			return err
		}
		appName := spec["metadata"].(map[interface{}]interface{})["name"].(string)
		syncCmd := []string{"argocd", "app", "sync", appName}
		if config.ArgoCD.Insecure {
			syncCmd = append(syncCmd, "--insecure")
		}
		if config.ArgoCD.GRPCWeb {
			syncCmd = append(syncCmd, "--grpc-web")
		}
		_, err = container.WithExec(syncCmd).Stdout(context.Background())
		return err
	},
}

func init() {
	engine.RegisterPatternMatches([]engine.PatternMatch{
		{Kind: engine.FilePatternMatch, Pattern: ".*.yaml"},
	})
	engine.RegisterAdmissionResolver(
		plan.ActionSpec{
			Name:        argocdSync.Name,
			DisplayName: "ArgoCD Sync",
			Description: "Sync the ArgoCD application with the source repo.",
		},
		[]engine.Fact{rules.ArgoCDApplicationExistsFact},
		nil,
		[]string{
			string(plan.ArgoCDServer),
			string(plan.ArgoCDAuthToken),
		},
	)
	plan.RegisterAction(argocdSync)
}
