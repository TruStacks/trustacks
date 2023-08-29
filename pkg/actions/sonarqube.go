package actions

import (
	"context"

	"dagger.io/dagger"
	"github.com/mitchellh/mapstructure"
	"github.com/trustacks/pkg/engine"
	"github.com/trustacks/pkg/engine/rules"
	"github.com/trustacks/pkg/plan"
)

var sonarScannerCLIScan = &plan.Action{
	Name:  "sonarScannerCLIScan",
	Image: "sonarsource/sonar-scanner-cli",
	Stage: plan.FeedbackStage,
	Script: func(container *dagger.Container, inputs map[string]interface{}, utils *plan.ActionUtilities) error {
		args := struct {
			SonarqubeToken string
		}{}
		if err := mapstructure.Decode(inputs, &args); err != nil {
			return err
		}
		_, err := container.
			WithSecretVariable("SONAR_TOKEN", utils.SetSecret("sonarqubeToken", args.SonarqubeToken)).
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
		[]engine.Fact{
			rules.SonarProjectPropertiesExists,
		},
		[]string{
			string(plan.SonarqubeToken),
		},
	)
	plan.RegisterAction(sonarScannerCLIScan)
}
