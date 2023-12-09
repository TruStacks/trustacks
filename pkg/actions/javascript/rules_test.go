package javascript

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPackageJSONExistsRule(t *testing.T) {
	t.Run("PackageJSONExistsFact is true", func(t *testing.T) {
		d, err := os.MkdirTemp("", "test-src")
		if err != nil {
			t.Fatal(err)
		}
		defer os.RemoveAll(d)
		if err := os.WriteFile(filepath.Join(d, "package.json"), []byte(`{}`), 0744); err != nil {
			t.Fatal(err)
		}
		fact, err := PackageJSONExistsRule(d, nil, nil)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, fact, PackageJSONExistsFact)
	})

	t.Run("PackageJSONExistsFact is false", func(t *testing.T) {
		d, err := os.MkdirTemp("", "test-src")
		if err != nil {
			t.Fatal(err)
		}
		defer os.RemoveAll(d)
		fact, err := PackageJSONExistsRule(d, nil, nil)
		if err != nil {
			t.Fatal(err)
		}
		assert.NotEqual(t, fact, PackageJSONExistsFact)
	})
}

func TestPackageJSONVersionExistsRule(t *testing.T) {
	t.Run("PackageJSONVersionExistsFact is true", func(t *testing.T) {
		d, err := os.MkdirTemp("", "test-src")
		if err != nil {
			t.Fatal(err)
		}
		defer os.RemoveAll(d)
		if err := os.WriteFile(filepath.Join(d, "package.json"), []byte(`{"version": "0.0.42"}`), 0744); err != nil {
			t.Fatal(err)
		}
		fact, err := PackageJSONVersionExistsRule(d, nil, nil)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, fact, PackageJSONVersionExistsFact)
	})

	t.Run("PackageJSONVersionExistsFact is false", func(t *testing.T) {
		d, err := os.MkdirTemp("", "test-src")
		if err != nil {
			t.Fatal(err)
		}
		defer os.RemoveAll(d)
		if err := os.WriteFile(filepath.Join(d, "package.json"), []byte(`{}`), 0744); err != nil {
			t.Fatal(err)
		}
		fact, err := PackageJSONVersionExistsRule(d, nil, nil)
		if err != nil {
			t.Fatal(err)
		}
		assert.NotEqual(t, fact, PackageJSONVersionExistsFact)
	})
}
