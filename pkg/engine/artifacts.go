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
	BuildArtifact Artifact = iota
	SemanticVersionArtifact
	ContainerImageArtifact
	CoverageArtifact
)

type Artifact int

type ArtifactStore struct {
	client    *dagger.Client
	artifacts map[Artifact]*dagger.Container
	mounts    []*ArtifactMount
}

var ErrArtifactNotFound = errors.New("artifact does not exists")

func (as *ArtifactStore) Mount(container *dagger.Container, artifact Artifact) (*dagger.Container, *ArtifactMount, error) {
	if _, ok := as.artifacts[artifact]; !ok {
		return container, nil, ErrArtifactNotFound
	}
	mount, err := newArtifactMount(artifact)
	if err != nil {
		return container, nil, err
	}
	as.mounts = append(as.mounts, mount)
	if _, err = as.artifacts[artifact].Directory(as.artifactPath(artifact)).Export(context.Background(), mount.hostDir); err != nil {
		return container, nil, err
	}
	return container.WithDirectory(mount.path, as.client.Host().Directory(mount.hostDir)), mount, nil
}

func (as *ArtifactStore) MountImage(container *dagger.Container, artifact Artifact) (*dagger.Container, *ArtifactMount, error) {
	if _, ok := as.artifacts[artifact]; !ok {
		return container, nil, ErrArtifactNotFound
	}
	mount, err := newArtifactMount(artifact)
	if err != nil {
		return container, nil, err
	}
	as.mounts = append(as.mounts, mount)
	if _, err = as.artifacts[artifact].Export(context.Background(), filepath.Join(mount.hostDir, "image.tar")); err != nil {
		return container, nil, err
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

type ArtifactMount struct {
	hostDir string
	path    string
}

func (m *ArtifactMount) Path(subpath string) string {
	if subpath == "" {
		return m.path
	}
	return filepath.Join(m.path, subpath)
}

func (m *ArtifactMount) Close() {
	defer os.RemoveAll(m.hostDir)
}

func newArtifactMount(artifact Artifact) (*ArtifactMount, error) {
	d, err := os.MkdirTemp("", fmt.Sprintf("%d-artifact", artifact))
	if err != nil {
		return nil, err
	}
	mount := &ArtifactMount{hostDir: d, path: filepath.Join("/tmp", shortuuid.New())}
	return mount, nil
}
