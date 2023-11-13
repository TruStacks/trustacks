package engine

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"dagger.io/dagger"
	"github.com/lithammer/shortuuid"
)

const (
	ApplicationDistArtifact Artifact = iota
	ContainerImageArtifact
	SemanticVersionArtifact
)

type Artifact int

type ArtifactStore struct {
	client    *dagger.Client
	artifacts map[Artifact]*dagger.Container
	mounts    []*artifactMount
}

func (as *ArtifactStore) Mount(container *dagger.Container, artifact Artifact) (*dagger.Container, *artifactMount, error) {
	if _, ok := as.artifacts[artifact]; !ok {
		return nil, nil, errors.New("artifact does not exists")
	}
	mount, err := newArtifactMount(artifact)
	if err != nil {
		return nil, nil, err
	}
	as.mounts = append(as.mounts, mount)
	if _, err = as.artifacts[artifact].Directory(as.artifactPath(artifact)).Export(context.Background(), mount.hostDir); err != nil {
		return nil, nil, err
	}
	return container.WithDirectory(mount.path, as.client.Host().Directory(mount.hostDir)), mount, nil
}

func (as *ArtifactStore) MountImage(container *dagger.Container, artifact Artifact) (*dagger.Container, *artifactMount, error) {
	if _, ok := as.artifacts[artifact]; !ok {
		return nil, nil, errors.New("artifact does not exists")
	}
	mount, err := newArtifactMount(artifact)
	if err != nil {
		return nil, nil, err
	}
	as.mounts = append(as.mounts, mount)
	if _, err = as.artifacts[artifact].Export(context.Background(), filepath.Join(mount.hostDir, "image.tar")); err != nil {
		return nil, nil, err
	}
	return container.WithDirectory(mount.path, as.client.Host().Directory(mount.hostDir)), mount, nil
}

func (as *ArtifactStore) Export(container *dagger.Container, artifact Artifact, path string) error {
	if _, ok := as.artifacts[artifact]; ok {
		return fmt.Errorf("artifact with id '%d' already exists", artifact)
	}
	container = container.WithExec([]string{"mkdir", "-p", as.artifactPath(artifact)})
	container = container.WithExec([]string{"mv", path, as.artifactPath(artifact)})
	as.artifacts[artifact] = container
	return nil
}

func (as *ArtifactStore) ExportContainer(container *dagger.Container, artifact Artifact) error {
	if _, ok := as.artifacts[artifact]; ok {
		return fmt.Errorf("artifact with id '%d' already exists", artifact)
	}
	as.artifacts[artifact] = container
	return nil
}

func (as *ArtifactStore) artifactPath(artifact Artifact) string {
	return fmt.Sprintf("/tmp/_artifacts/%d", artifact)
}

func newArtifactStore(client *dagger.Client) *ArtifactStore {
	return &ArtifactStore{
		client:    client,
		artifacts: make(map[Artifact]*dagger.Container),
	}
}

type artifactMount struct {
	hostDir string
	path    string
}

func (m *artifactMount) Path(subpath string) string {
	if subpath == "" {
		return m.path
	}
	return filepath.Join(m.path, subpath)
}

func (m *artifactMount) Close() {
	defer os.RemoveAll(m.hostDir)
}

func newArtifactMount(artifact Artifact) (*artifactMount, error) {
	d, err := os.MkdirTemp("", fmt.Sprintf("%d-artifact", artifact))
	if err != nil {
		return nil, err
	}
	mount := &artifactMount{hostDir: d, path: shortuuid.New()}
	return mount, nil
}
