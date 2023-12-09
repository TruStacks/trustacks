package argocd

import (
	"context"
	"strings"

	"dagger.io/dagger"
	"github.com/mitchellh/mapstructure"
	"github.com/trustacks/trustacks/pkg/engine"
	"gopkg.in/yaml.v2"
)

// extraOptions returns optional global command options dependent on
// parameters in the declarative configuration.
func extraGlobalOptions(config *engine.Config) []string {
	args := []string{}
	if config.ArgoCD.Insecure {
		args = append(args, "--insecure")
	}
	if config.ArgoCD.GRPCWeb {
		args = append(args, "--grpc-web")
	}
	return args
}

// getArgoApplicationInfo gets the argo application manifest path and
// the application metadata name from the manifest.
func getArgoApplicationInfo(container *dagger.Container) (string, string, error) {
	container = container.WithExec([]string{"grep", "-r", "argoproj.io/v1alpha1"})
	stdout, err := container.Stdout(context.Background())
	if err != nil {
		return "", "", err
	}
	appSpecPath := strings.Split(strings.Split(stdout, "\n")[0], ":")[0]
	container = container.WithExec([]string{"cat", appSpecPath})
	stdout, err = container.Stdout(context.Background())
	if err != nil {
		return "", "", err
	}
	var spec ApplicationSpec
	if err := yaml.Unmarshal([]byte(stdout), &spec); err != nil {
		return "", "", err
	}
	return appSpecPath, spec.Metadata.Name, nil
}

// argocdSync is an action that creates and syncs an argo cd
// application to a kubernetes cluster.
var argocdSync = &engine.Action{
	Name:        "argocdSync",
	DisplayName: "Argo CD Sync",
	Description: "Sync the Argo CD application with the source repo.",
	Image:       func(_ *engine.Config) string { return "argoproj/argocd" },
	Stage:       engine.ReleaseStage,
	Script: func(container *dagger.Container, inputs map[string]interface{}, utils *engine.ActionUtilities) error {
		var err error
		args := struct {
			ARGOCD_SERVER     string //nolint:revive,stylecheck
			ARGOCD_AUTH_TOKEN string //nolint:revive,stylecheck
		}{}
		if err := mapstructure.Decode(inputs, &args); err != nil {
			return err
		}
		appSpecPath, appName, err := getArgoApplicationInfo(container)
		if err != nil {
			return err
		}
		extraOpts := extraGlobalOptions(utils.GetConfig())
		container = container.WithSecretVariable("ARGOCD_AUTH_TOKEN", utils.SetSecret("argocdAuthToken", args.ARGOCD_AUTH_TOKEN))
		container = container.WithEnvVariable("ARGOCD_SERVER", args.ARGOCD_SERVER)
		container = container.WithExec(append([]string{"argocd", "app", "create", "-f", appSpecPath, "--upsert"}, extraOpts...))
		_, err = container.WithExec(append([]string{"argocd", "app", "sync", appName}, extraOpts...)).Sync(context.Background())
		return err
	},
	Inputs: []engine.InputField{
		engine.ArgoCDServer,
		engine.ArgoCDAuthToken,
	},
	AdmissionCriteria: []engine.Fact{ArgoCDApplicationExistsFact},
}

func init() {
	engine.RegisterPatternMatches([]engine.PatternMatch{
		{
			Kind:    engine.FilePatternMatch,
			Pattern: ".*.yaml",
		},
	})
	engine.RegisterAction(argocdSync)
}
