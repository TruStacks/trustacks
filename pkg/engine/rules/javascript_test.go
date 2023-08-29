package rules

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPackageJsonExistsRule(t *testing.T) {
	t.Run("PackageJsonExistsFact is true", func(t *testing.T) {
		d, err := os.MkdirTemp("", "test-src")
		if err != nil {
			t.Fatal(err)
		}
		defer os.RemoveAll(d)
		if err := os.WriteFile(filepath.Join(d, "package.json"), []byte(`{}`), 0744); err != nil {
			t.Fatal(err)
		}
		fact, err := packageJsonExistsRule(d, nil, nil)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, fact, PackageJsonExistsFact)
	})

	t.Run("PackageJsonExistsFact is false", func(t *testing.T) {
		d, err := os.MkdirTemp("", "test-src")
		if err != nil {
			t.Fatal(err)
		}
		defer os.RemoveAll(d)
		fact, err := packageJsonExistsRule(d, nil, nil)
		if err != nil {
			t.Fatal(err)
		}
		assert.NotEqual(t, fact, PackageJsonExistsFact)
	})
}

func TestPackageJsonVersionExistsRule(t *testing.T) {
	t.Run("PackageJsonVersionExistsFact is true", func(t *testing.T) {
		d, err := os.MkdirTemp("", "test-src")
		if err != nil {
			t.Fatal(err)
		}
		defer os.RemoveAll(d)
		if err := os.WriteFile(filepath.Join(d, "package.json"), []byte(`{"version": "0.0.42"}`), 0744); err != nil {
			t.Fatal(err)
		}
		fact, err := packageJsonVersionExistsRule(d, nil, nil)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, fact, PackageJsonVersionExistsFact)
	})

	t.Run("PackageJsonVersionExistsFact is false", func(t *testing.T) {
		d, err := os.MkdirTemp("", "test-src")
		if err != nil {
			t.Fatal(err)
		}
		defer os.RemoveAll(d)
		if err := os.WriteFile(filepath.Join(d, "package.json"), []byte(`{}`), 0744); err != nil {
			t.Fatal(err)
		}
		fact, err := packageJsonVersionExistsRule(d, nil, nil)
		if err != nil {
			t.Fatal(err)
		}
		assert.NotEqual(t, fact, PackageJsonVersionExistsFact)
	})
}
