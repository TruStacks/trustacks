package container

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainerfileExistsRule(t *testing.T) {
	t.Run("ContainerfileExistFact is true", func(t *testing.T) {
		tests := []struct {
			filename string
		}{
			{"Dockerfile"},
			{"Containerfile"},
		}
		tempDirs := []string{}
		defer func() {
			for _, d := range tempDirs {
				os.RemoveAll(d)
			}
		}()
		for _, tc := range tests {
			d, err := os.MkdirTemp("", "test-src")
			if err != nil {
				t.Fatal(err)
			}
			tempDirs = append(tempDirs, d)
			if err := os.WriteFile(filepath.Join(d, tc.filename), []byte("FROM alpine"), 0744); err != nil {
				t.Fatal(err)
			}
			fact, err := ContainerfileExistsRule(d, nil, nil)
			if err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, fact, ContainerfileExistFact)
		}
	})

	t.Run("ContainerfileExistFact is false", func(t *testing.T) {
		d, err := os.MkdirTemp("", "test-src")
		if err != nil {
			t.Fatal(err)
		}
		fact, err := ContainerfileExistsRule(d, nil, nil)
		if err != nil {
			t.Fatal(err)
		}
		assert.NotEqual(t, fact, ContainerfileExistFact)
	})
}

func TestContainerfileHasNoDependenciesRule(t *testing.T) {
	t.Run("ContainerfileHasNoDependenciesRule is true", func(t *testing.T) {
		tests := []struct {
			filename string
		}{
			{"Dockerfile"},
			{"Containerfile"},
		}
		tempDirs := []string{}
		defer func() {
			for _, d := range tempDirs {
				os.RemoveAll(d)
			}
		}()
		for _, tc := range tests {
			d, err := os.MkdirTemp("", "test-src")
			if err != nil {
				t.Fatal(err)
			}
			tempDirs = append(tempDirs, d)
			contents := []byte(`FROM alpine
COPY test.txt /tmp/test.txt`)
			if err := os.WriteFile(filepath.Join(d, tc.filename), contents, 0744); err != nil {
				t.Fatal(err)
			}
			if err := os.WriteFile(filepath.Join(d, "test.txt"), []byte("test"), 0744); err != nil {
				t.Fatal(err)
			}
			fact, err := ContainerfileHasNoDependenciesRule(d, nil, nil)
			if err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, fact, ContainerfileHasNoDependenciesFact)
		}
	})

	t.Run("containerfileHasNoDependenciesRule is false", func(t *testing.T) {
		d, err := os.MkdirTemp("", "test-src")
		if err != nil {
			t.Fatal(err)
		}
		contents := []byte(`FROM alpine
COPY test.txt /tmp/test.txt`)
		if err := os.WriteFile(filepath.Join(d, "Dockerfile"), contents, 0744); err != nil {
			t.Fatal(err)
		}
		fact, err := ContainerfileHasNoDependenciesRule(d, nil, nil)
		if err != nil {
			t.Fatal(err)
		}
		assert.NotEqual(t, fact, ContainerfileHasNoDependenciesFact)
	})
}
