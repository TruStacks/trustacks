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
	Stage:       engine.OnDemand,
	OutputArtifacts: []engine.Artifact{
		engine.ContainerImageArtifact,
	},
	OptionalInputArtifacts: []engine.Artifact{
		engine.BuildArtifact,
	},
	Script: func(container *dagger.Container, _ map[string]interface{}, utils *engine.ActionUtilities) error {
		container, buildMount, err := utils.Mount(container, engine.BuildArtifact)
		if err != nil && err != engine.ErrArtifactNotFound {
			return err
		} else if err == nil {
			container = container.WithExec([]string{"cp", "-r", buildMount.Path(".build"), "./.build"})
		}
		container = container.Directory("/src").DockerBuild()
		if err := utils.ExportContainer(container, engine.ContainerImageArtifact); err != nil {
			return err
		}
		_, err = container.Sync(context.Background())
		return err
	},
	AdmissionCriteria: []engine.Fact{ContainerfileHasPredictableDependenciesFact},
}

var containerPublishAction = &engine.Action{
	Name:        "containerPublish",
	DisplayName: "Container Publish",
	Description: "Publish the container to a container registry.",
	Image:       func(_ *engine.Config) string { return "alpine" },
	Stage:       engine.DeployStage,
	InputArtifacts: []engine.Artifact{
		engine.ContainerImageArtifact,
	},
	OptionalInputArtifacts: []engine.Artifact{
		engine.SemanticVersionArtifact,
	},
	Script: func(container *dagger.Container, inputs map[string]interface{}, utils *engine.ActionUtilities) error {
		args := struct {
			CONTAINER_REGISTRY          string //nolint:revive,stylecheck
			CONTAINER_REGISTRY_USERNAME string //nolint:revive,stylecheck
			CONTAINER_REGISTRY_PASSWORD string //nolint:revive,stylecheck
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
		container = container.Import(container.File(imageMount.Path("image.tar")))
		container = container.WithRegistryAuth(
			args.CONTAINER_REGISTRY,
			args.CONTAINER_REGISTRY_USERNAME,
			utils.SetSecret("registryPassword", args.CONTAINER_REGISTRY_PASSWORD),
		)
		_, err = container.Publish(
			context.Background(),
			fmt.Sprintf("%s:%s", args.CONTAINER_REGISTRY, strings.ReplaceAll(version, "\n", "")),
		)
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
	AdmissionCriteria: []engine.Fact{ContainerfileHasPredictableDependenciesFact},
}

func init() {
	engine.RegisterAction(containerBuildAction)
	engine.RegisterAction(containerPublishAction)
}
