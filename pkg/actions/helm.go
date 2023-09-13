package actions

import (
	"context"
	"io/fs"
	"os"
	"path"
	"path/filepath"

	"dagger.io/dagger"
	"github.com/mitchellh/mapstructure"
	"github.com/trustacks/trustacks/pkg/engine"
	"github.com/trustacks/trustacks/pkg/engine/rules"
	"github.com/trustacks/trustacks/pkg/plan"
	"gopkg.in/yaml.v2"
)

var helmInstallStageAction = &plan.Action{
	Name:  "helmInstall[Stage]",
	Image: "alpine/helm",
	Stage: plan.StageStage,
	Script: func(container *dagger.Container, inputs map[string]interface{}, utils *plan.ActionUtilities) error {
		args := struct {
			KubernetesStagingKubeconfig string
			KubernetesNamespace         string
		}{}
		if err := mapstructure.Decode(inputs, &args); err != nil {
			return err
		}
		var chartPath string
		if err := filepath.WalkDir("./", func(path string, info fs.DirEntry, err error) error {
			if info.Name() == "Chart.yaml" {
				chartPath = path
			}
			return nil
		}); err != nil {
			return err
		}
		chart := map[string]interface{}{}
		contents, err := os.ReadFile(chartPath)
		if err != nil {
			return err
		}
		if yaml.Unmarshal(contents, &chart); err != nil {
			return err
		}
		container = container.WithNewFile("/tmp/kubeconfig", dagger.ContainerWithNewFileOpts{Contents: args.KubernetesStagingKubeconfig, Permissions: 0600})
		container = container.WithSecretVariable("KUBECONFIG", utils.SetSecret("kubeconfig", "/tmp/kubeconfig"))
		container = container.WithExec([]string{"upgrade", chart["name"].(string), "--install", "--namespace", args.KubernetesNamespace, "--wait", "--wait-for-jobs", path.Dir(chartPath)})
		_, err = container.Stdout(context.Background())
		return err
	},
}

func init() {
	engine.RegisterPatternMatches([]engine.PatternMatch{
		{
			Kind:    engine.FilePatternMatch,
			Pattern: "Chart.yaml",
		},
	})
	engine.RegisterAdmissionResolver(
		plan.ActionSpec{
			Name:        helmInstallStageAction.Name,
			DisplayName: "Helm Install [Stage]",
			Description: "Install the application into a kubernetes staging cluster with helm.",
		},
		[]engine.Fact{rules.HelmChartExistsFact},
		[]string{
			string(plan.KubernetesStagingKubeconfig),
			string(plan.KubernetesNamespace),
		})
	plan.RegisterAction(helmInstallStageAction)
}
