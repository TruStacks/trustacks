package actions

import (
	"context"

	"dagger.io/dagger"
	"github.com/trustacks/pkg/engine"
	"github.com/trustacks/pkg/engine/rules"
	"github.com/trustacks/pkg/plan"
)

var trivyImageAction = &plan.Action{
	Name:   "trivyImageAction",
	Image:  "aquasec/trivy",
	Stage:  plan.FeedbackStage,
	Caches: []string{"/src/node_modules"},
	InputArtifacts: []plan.Artifact{
		plan.ContainerImageArtifact,
	},
	Script: func(container *dagger.Container, _ map[string]interface{}, utils *plan.ActionUtilities) error {
		container, imageMount, err := utils.MountImage(container, plan.ContainerImageArtifact)
		if err != nil {
			return err
		}
		container = container.WithExec([]string{"image", "--input", imageMount.Path("image.tar")})
		_, err = container.Stdout(context.Background())
		return err
	},
}

func init() {
	engine.RegisterAdmissionResolver(
		plan.ActionSpec{
			Name:        trivyImageAction.Name,
			DisplayName: "Trivy Scan",
			Description: "Scan the container image with the trivy security scanner.",
		},
		[]engine.Fact{rules.TrivyConfigExistsFact},
		nil,
	)
	plan.RegisterAction(trivyImageAction)
}
