package container

import (
	"context"
	"fmt"
	"strings"

	"dagger.io/dagger"
	"github.com/mitchellh/mapstructure"
	"github.com/trustacks/trustacks/pkg/engine"
)

var containerBuildAction = &engine.Action{
	Name:        "containerBuild",
	DisplayName: "Container Build",
	Description: "Build a container image from the source Containerfile or Dockerfile.",
	Image:       func(_ *engine.Config) string { return "busybox" },
	Stage:       engine.OnDemandStage,
	OutputArtifacts: []engine.Artifact{
		engine.ContainerImageArtifact,
	},
	Script: func(container *dagger.Container, _ map[string]interface{}, utils *engine.ActionUtilities) error {
		container = container.Directory("/src").DockerBuild()
		return utils.ExportContainer(container, engine.ContainerImageArtifact)
	},
	AdmissionCriteria: []engine.Fact{ContainerfileHasNoDependenciesFact},
}

var containerCopyAction = &engine.Action{
	Name:        "containerCopy",
	DisplayName: "Container Publish",
	Description: "Publish the container to a container registry.",
	Image:       func(_ *engine.Config) string { return "alpine" },
	Stage:       engine.PackageStage,
	InputArtifacts: []engine.Artifact{
		engine.ContainerImageArtifact,
		engine.SemanticVersionArtifact,
	},
	Script: func(container *dagger.Container, inputs map[string]interface{}, utils *engine.ActionUtilities) error {
		args := struct {
			CONTAINER_REGISTRY          string
			CONTAINER_REGISTRY_USERNAME string
			CONTAINER_REGISTRY_PASSWORD string
		}{}
		if err := mapstructure.Decode(inputs, &args); err != nil {
			return err
		}
		container, imageMount, err := utils.MountImage(container, engine.ContainerImageArtifact)
		if err != nil {
			return err
		}
		version := utils.GetConfig().Common.Version
		if version == "" {
			container, versionMount, err := utils.Mount(container, engine.SemanticVersionArtifact)
			if err != nil {
				return err
			}
			version, err = container.File(versionMount.Path("version")).Contents(context.Background())
			if err != nil {
				return err
			}
		}
		_, err = container.Import(container.File(imageMount.Path("image.tar"))).
			WithRegistryAuth(args.CONTAINER_REGISTRY, args.CONTAINER_REGISTRY_USERNAME, utils.SetSecret("registryPassword", args.CONTAINER_REGISTRY_PASSWORD)).
			Publish(context.Background(), fmt.Sprintf("%s:%s", args.CONTAINER_REGISTRY, strings.ReplaceAll(version, "\n", "")))
		if err != nil {
			return err
		}
		return err
	},
	Inputs: []engine.InputField{
		engine.ContainerRegistry,
		engine.ContainerRegistryUsername,
		engine.ContainerRegistryPassword,
	},
	AdmissionCriteria: []engine.Fact{ContainerfileHasNoDependenciesFact},
}

func init() {
	engine.RegisterAction(containerBuildAction)
	engine.RegisterAction(containerCopyAction)
}
