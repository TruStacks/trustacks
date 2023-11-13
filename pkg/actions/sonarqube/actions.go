package sonarqube

import (
	"context"

	"dagger.io/dagger"
	"github.com/mitchellh/mapstructure"
	"github.com/trustacks/trustacks/pkg/engine"
)

var sonarScannerCLIScan = &engine.Action{
	Name:        "sonarScannerCLIScan",
	DisplayName: "SonarQube Scan",
	Description: "Scan the source with the sonar scanner cli.",
	Image:       func(_ *engine.Config) string { return "sonarsource/sonar-scanner-cli" },
	Stage:       engine.FeedbackStage,
	Script: func(container *dagger.Container, inputs map[string]interface{}, utils *engine.ActionUtilities) error {
		args := struct {
			SONARQUBE_TOKEN string
		}{}
		if err := mapstructure.Decode(inputs, &args); err != nil {
			return err
		}
		_, err := container.
			WithSecretVariable("SONAR_TOKEN", utils.SetSecret("sonarqubeToken", args.SONARQUBE_TOKEN)).
			Stdout(context.Background())
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
