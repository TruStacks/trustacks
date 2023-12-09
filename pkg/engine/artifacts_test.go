package engine

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"dagger.io/dagger"
	"github.com/stretchr/testify/assert"
)

func TestMountIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	var mockArtifact Artifact = 23
	client, err := dagger.Connect(context.Background(), dagger.WithLogOutput(os.Stdout))
	if err != nil {
		t.Fatal(err)
	}
	store := newArtifactStore(client)
	container := client.Container().
		From("alpine").
		WithNewFile(
			filepath.Join(store.artifactPath(mockArtifact), "hello"),
			dagger.ContainerWithNewFileOpts{Contents: "Hello, World!"},
		)
	store.artifacts[mockArtifact] = container
	container, mockMount, err := store.Mount(container, mockArtifact)
	if err != nil {
		t.Fatal(err)
	}
	contents, err := container.File(mockMount.Path("hello")).Contents(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "Hello, World!", contents)
}

func TestMountImageIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	var mockImageArtifact Artifact = 23
	client, err := dagger.Connect(context.Background(), dagger.WithLogOutput(os.Stdout))
	if err != nil {
		t.Fatal(err)
	}
	store := newArtifactStore(client)
	container := client.Container().From("alpine")
	store.artifacts[mockImageArtifact] = container
	container, mockMount, err := store.MountImage(container, mockImageArtifact)
	if err != nil {
		t.Fatal(err)
	}
	_, err = container.File(mockMount.Path("image.tar")).Contents(context.Background())
	if err != nil {
		t.Fatal(err)
	}
}

func TestExportIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	var mockArtifact Artifact = 23
	client, err := dagger.Connect(context.Background(), dagger.WithLogOutput(os.Stdout))
	if err != nil {
		t.Fatal(err)
	}
	store := newArtifactStore(client)
	container := client.Container().
		From("alpine").
		WithNewFile(filepath.Join("/tmp", "hello"), dagger.ContainerWithNewFileOpts{Contents: "Hello, World!"})
	if err := store.Export(container, mockArtifact, filepath.Join("/tmp", "hello")); err != nil {
		t.Fatal(err)
	}
	container, mockMount, err := store.Mount(container, mockArtifact)
	if err != nil {
		t.Fatal(err)
	}
	contents, err := container.File(mockMount.Path("hello")).Contents(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "Hello, World!", contents)
}

func TestExportImageIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	var mockImageArtifact Artifact = 23
	client, err := dagger.Connect(context.Background(), dagger.WithLogOutput(os.Stdout))
	if err != nil {
		t.Fatal(err)
	}
	store := newArtifactStore(client)
	container := client.Container().From("alpine")
	if err := store.ExportContainer(container, mockImageArtifact); err != nil {
		t.Fatal(err)
	}
	container, mockMount, err := store.MountImage(container, mockImageArtifact)
	if err != nil {
		t.Fatal(err)
	}
	_, err = container.File(mockMount.Path("image.tar")).Contents(context.Background())
	if err != nil {
		t.Fatal(err)
	}
}

func TestArtifactPath(t *testing.T) {
	var mockArtifact Artifact = 23
	store := &ArtifactStore{}
	assert.Equal(t, "/tmp/_artifacts/23", store.artifactPath(mockArtifact))
}
