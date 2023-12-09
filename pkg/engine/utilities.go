package engine

import (
	"context"

	"dagger.io/dagger"
)

const (
	DockerCLIOnDebian = "debian"
)

type ActionUtilities struct {
	*ArtifactStore
	client *dagger.Client
	config *Config
}

func (util *ActionUtilities) SetSecret(name, plaintext string) *dagger.Secret {
	return util.client.SetSecret(name, plaintext)
}

func (util *ActionUtilities) GetConfig() *Config {
	return util.config
}

func (util *ActionUtilities) WithDockerdService(container *dagger.Container) (*dagger.Container, func(), error) {
	dockerdPort := 2376
	dockerClientCerts := util.client.CacheVolume("trustacks-docker-client-certs")
	dockerd, err := util.client.Container().
		From("docker:24.0.7-dind-rootless").
		WithMountedCache("/certs/client", dockerClientCerts).
		WithExec(nil, dagger.ContainerWithExecOpts{InsecureRootCapabilities: true}).
		WithoutExposedPort(dockerdPort).
		AsService().
		Start(context.Background())
	if err != nil {
		return container, nil, err
	}
	container = container.
		WithServiceBinding("docker", dockerd).
		WithMountedCache("/tmp/docker/client", dockerClientCerts).
		WithEnvVariable("DOCKER_CERT_PATH", "/tmp/docker/client").
		WithEnvVariable("DOCKER_TLS_VERIFY", "1").
		WithEnvVariable("DOCKER_HOST", "tcp://docker:2376")
	return container, func() {
		dockerd.Stop(context.Background()) //nolint:errcheck
	}, nil
}

func (util *ActionUtilities) WithDockerCLI(distro string, container *dagger.Container) *dagger.Container {
	//@TODO: remove nolint after adding additional distros.
	//nolint:gocritic
	switch distro {
	case DockerCLIOnDebian:
		// install docker on debian:
		// https://docs.docker.com/engine/install/debian/
		container = container.WithExec([]string{"/bin/sh", "-c", `
apt-get update
apt-get install -y ca-certificates curl gnupg
install -m 0755 -d /etc/apt/keyrings
curl -fsSL https://download.docker.com/linux/debian/gpg | gpg --dearmor -o /etc/apt/keyrings/docker.gpg
chmod a+r /etc/apt/keyrings/docker.gpg
echo \
"deb [arch="$(dpkg --print-architecture)" signed-by=/etc/apt/keyrings/docker.gpg] https://download.docker.com/linux/debian \
"$(. /etc/os-release && echo "$VERSION_CODENAME")" stable" | \
tee /etc/apt/sources.list.d/docker.list > /dev/null
apt-get update
apt install
apt-get install -y docker-ce docker-ce-cli containerd.io docker-buildx-plugin
		`})
	}
	return container
}

func newActionUtilities(client *dagger.Client, artifacts *ArtifactStore, config *Config) *ActionUtilities {
	return &ActionUtilities{artifacts, client, config}
}
