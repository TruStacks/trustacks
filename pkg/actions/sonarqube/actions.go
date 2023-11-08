package sonarqube

import (
	"context"

	"dagger.io/dagger"
	"github.com/mitchellh/mapstructure"
	"github.com/trustacks/trustacks/pkg/engine"
	"github.com/trustacks/trustacks/pkg/plan"
)

var sonarScannerCLIScan = &plan.Action{
	Name:  "sonarScannerCLIScan",
	Image: func(_ *plan.Config) string { return "sonarsource/sonar-scanner-cli" },
	Stage: plan.FeedbackStage,
	Script: func(container *dagger.Container, inputs map[string]interface{}, utils *plan.ActionUtilities) error {
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
}

func init() {
	engine.RegisterAdmissionResolver(
		plan.ActionSpec{
			Name:        sonarScannerCLIScan.Name,
			DisplayName: "SonarQube Scan",
			Description: "Scan the source with the sonar scanner cli.",
		},
		[]engine.Fact{SonarProjectPropertiesExists},
		nil,
		[]string{
			string(plan.SonarqubeToken),
		},
	)
	plan.RegisterAction(sonarScannerCLIScan)
}
