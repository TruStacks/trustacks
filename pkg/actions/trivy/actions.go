package trivy

import (
	"context"

	"dagger.io/dagger"
	"github.com/trustacks/trustacks/pkg/engine"
)

var trivyImageAction = &engine.Action{
	Name:        "trivyImage",
	DisplayName: "Trivy Scan",
	Description: "Scan the container image with the trivy security scanner.",
	Image:       func(_ *engine.Config) string { return "aquasec/trivy" },
	Stage:       engine.FeedbackStage,
	Caches:      []string{"/src/node_modules"},
	InputArtifacts: []engine.Artifact{
		engine.ContainerImageArtifact,
	},
	Script: func(container *dagger.Container, _ map[string]interface{}, utils *engine.ActionUtilities) error {
		container, imageMount, err := utils.MountImage(container, engine.ContainerImageArtifact)
		if err != nil {
			return err
		}
		container = container.WithExec([]string{"image", "--input", imageMount.Path("image.tar")})
		_, err = container.Stdout(context.Background())
		return err
	},
	AdmissionCriteria: []engine.Fact{TrivyConfigExistsFact},
}

func init() {
	engine.RegisterAction(trivyImageAction)
}
