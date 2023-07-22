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
	Image: "alpine",
	State: plan.OnDemandState,
	InputArtifacts: []plan.Artifact{
		plan.ApplicationDistArtifact,
	},
	OutputArtifacts: []plan.Artifact{
		plan.ContainerImageArtifact,
	},
	Script: func(container *dagger.Container, _ map[string]interface{}, utils *plan.ActionUtilities) error {
		container, appDistMount, err := utils.Mount(container, plan.ApplicationDistArtifact)
		if err != nil {
			return err
		}
		container = container.Directory("/src").DockerBuild(dagger.DirectoryDockerBuildOpts{
			BuildArgs: []dagger.BuildArg{
				{
					Name:  "app_dist",
					Value: appDistMount.Path("build"),
				},
			},
		})
		return utils.ExportContainer(container, plan.ContainerImageArtifact)
	},
}

var containerCopyAction = &plan.Action{
	Name:  "containerCopy",
	Image: "alpine",
	State: plan.StageState,
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
	engine.RegisterAdmissionResolver(containerBuildAction.Name, []engine.Fact{engine.ContainerfileExistFact}, nil)
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
