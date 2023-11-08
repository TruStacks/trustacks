package npm

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

type npmTestCollector struct {
	results []string
}

func (c npmTestCollector) Search(pattern string) []string {
	return c.results
}

func TestNpmTestExsitsRule(t *testing.T) {
	t.Run("NpmTestExistsFact is true", func(t *testing.T) {
		d, err := os.MkdirTemp("", "test-src")
		if err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(filepath.Join(d, "package.json"), []byte(`{"scripts":{"test": "react-scripts test"}}`), 0744); err != nil {
			t.Fatal(err)
		}
		collector := npmTestCollector{results: []string{"my.test.js"}}
		fact, err := NpmTestExsitsRule(d, collector, nil)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, fact, NpmTestExistsFact)
	})

	t.Run("NpmTestExistsFact is false", func(t *testing.T) {
		d, err := os.MkdirTemp("", "test-src")
		if err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(filepath.Join(d, "package.json"), []byte("{}"), 0744); err != nil {
			t.Fatal(err)
		}
		collector := npmTestCollector{results: []string{"my.test.js"}}
		fact, err := NpmTestExsitsRule(d, collector, nil)
		if err != nil {
			t.Fatal(err)
		}
		assert.NotEqual(t, fact, NpmTestExistsFact)
	})
}

func TestNpmBuildExsitsRule(t *testing.T) {
	t.Run("NpmBuildExistsFact is true", func(t *testing.T) {
		d, err := os.MkdirTemp("", "test-src")
		if err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(filepath.Join(d, "package.json"), []byte(`{"scripts":{"build": "react-scripts build"}}`), 0744); err != nil {
			t.Fatal(err)
		}
		fact, err := NpmBuildExsitsRule(d, nil, nil)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, fact, NpmBuildExistsFact)
	})

	t.Run("NpmBuildExistsFact is false", func(t *testing.T) {
		d, err := os.MkdirTemp("", "test-src")
		if err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(filepath.Join(d, "package.json"), []byte("{}"), 0744); err != nil {
			t.Fatal(err)
		}
		fact, err := NpmBuildExsitsRule(d, nil, nil)
		if err != nil {
			t.Fatal(err)
		}
		assert.NotEqual(t, fact, NpmBuildExistsFact)
	})
}
