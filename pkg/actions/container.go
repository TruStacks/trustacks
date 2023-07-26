package actions

import (
	"context"
	"fmt"
	"strings"

	"dagger.io/dagger"
	"github.com/mitchellh/mapstructure"
	"github.com/trustacks/pkg/engine"
	"github.com/trustacks/pkg/plan"
)

var containerBuildAction = &plan.Action{
	Name:  "containerBuild",
	Image: "busybox",
	State: plan.OnDemandState,
	OutputArtifacts: []plan.Artifact{
		plan.ContainerImageArtifact,
	},
	Script: func(container *dagger.Container, _ map[string]interface{}, utils *plan.ActionUtilities) error {
		container = container.Directory("/src").DockerBuild()
		return utils.ExportContainer(container, plan.ContainerImageArtifact)
	},
}

var containerCopyAction = &plan.Action{
	Name:  "containerCopy",
	Image: "alpine",
	State: plan.PackageState,
	InputArtifacts: []plan.Artifact{
		plan.SemanticVersionArtifact,
		plan.ContainerImageArtifact,
	},
	Script: func(container *dagger.Container, inputs map[string]interface{}, utils *plan.ActionUtilities) error {
		args := struct {
			ContainerRegistry         string
			ContainerRegistryUsername string
			ContainerRegistryPassword string
		}{}
		if err := mapstructure.Decode(inputs, &args); err != nil {
			return err
		}
		container, imageMount, err := utils.MountImage(container, plan.ContainerImageArtifact)
		if err != nil {
			return err
		}
		container, versionMount, err := utils.Mount(container, plan.SemanticVersionArtifact)
		if err != nil {
			return err
		}
		version, err := container.File(versionMount.Path("version")).Contents(context.Background())
		if err != nil {
			return err
		}
		_, err = container.Import(container.File(imageMount.Path("image.tar"))).
			WithRegistryAuth(args.ContainerRegistry, args.ContainerRegistryUsername, utils.SetSecret("registryPassword", args.ContainerRegistryPassword)).
			Publish(context.Background(), fmt.Sprintf("%s:%s", args.ContainerRegistry, strings.ReplaceAll(version, "\n", "")))
		if err != nil {
			return err
		}
		return err
	},
}

func init() {
	engine.RegisterAdmissionResolver(containerBuildAction.Name, []engine.Fact{engine.ContainerfileIsMultiStageFact}, nil)
	plan.RegisterAction(containerBuildAction)

	engine.RegisterAdmissionResolver(
		containerCopyAction.Name,
		[]engine.Fact{engine.ContainerfileExistFact},
		[]string{
			string(plan.ContainerRegistry),
			string(plan.ContainerRegistryUsername),
			string(plan.ContainerRegistryPassword),
		},
	)
	plan.RegisterAction(containerCopyAction)
}
