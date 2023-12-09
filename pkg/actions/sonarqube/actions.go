package sonarqube

import (
	"context"
	"fmt"

	"dagger.io/dagger"
	"github.com/mitchellh/mapstructure"
	"github.com/trustacks/trustacks/pkg/engine"
)

var sonarScannerCLIScan = &engine.Action{
	Name:        "sonarScannerCLIScan",
	DisplayName: "SonarQube Scan",
	Description: "Scan the source with the sonar scanner cli.",
	Image:       func(_ *engine.Config) string { return "sonarsource/sonar-scanner-cli" },
	Stage:       engine.CommitStage,
	OptionalInputArtifacts: []engine.Artifact{
		engine.CoverageArtifact,
	},
	Script: func(container *dagger.Container, inputs map[string]interface{}, utils *engine.ActionUtilities) error {
		args := struct {
			SONARQUBE_TOKEN string //nolint:revive,stylecheck
		}{}
		if err := mapstructure.Decode(inputs, &args); err != nil {
			return err
		}
		container, coverageMount, err := utils.Mount(container, engine.CoverageArtifact)
		if err != nil && err != engine.ErrArtifactNotFound {
			return err
		} else if err == nil {
			container = container.WithExec([]string{"/bin/sh", "-c", fmt.Sprintf("cp %s/* ./", coverageMount.Path(""))})
		}
		container = container.WithSecretVariable("SONAR_TOKEN", utils.SetSecret("sonarqubeToken", args.SONARQUBE_TOKEN))
		container = container.WithExec(nil)
		_, err = container.Sync(context.Background())
		return err
	},
	Inputs: []engine.InputField{
		engine.SonarqubeToken,
	},
	AdmissionCriteria: []engine.Fact{SonarProjectPropertiesExists},
}

func init() {
	engine.RegisterAction(sonarScannerCLIScan)
}
