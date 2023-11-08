package golang

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

type golangTestCollector struct {
	results []string
}

func (c golangTestCollector) Search(pattern string) []string {
	return c.results
}

func TestGoModExistsRule(t *testing.T) {
	t.Run("GoModExistsFact is true", func(t *testing.T) {
		d, err := os.MkdirTemp("", "test-src")
		if err != nil {
			t.Fatal(err)
		}
		defer os.RemoveAll(d)
		if err := os.WriteFile(filepath.Join(d, "go.mod"), []byte(``), 0744); err != nil {
			t.Fatal(err)
		}
		fact, err := GoModExistsRule(d, nil, nil)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, fact, GoModExistsFact)
	})

	t.Run("GoModExistsFact is false", func(t *testing.T) {
		d, err := os.MkdirTemp("", "test-src")
		if err != nil {
			t.Fatal(err)
		}
		defer os.RemoveAll(d)
		fact, err := GoModExistsRule(d, nil, nil)
		if err != nil {
			t.Fatal(err)
		}
		assert.NotEqual(t, fact, GoModExistsFact)
	})
}

func TestGolangTestsExistRule(t *testing.T) {
	t.Run("GolangTestsExistFact is true", func(t *testing.T) {
		d, err := os.MkdirTemp("", "test-src")
		if err != nil {
			t.Fatal(err)
		}
		defer os.RemoveAll(d)
		collector := golangTestCollector{results: []string{"my_test.go"}}
		fact, err := GolangTestExistsRule(d, collector, nil)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, fact, GolangTestsExistFact)
	})

	t.Run("GolangTestsExistFact is false", func(t *testing.T) {
		d, err := os.MkdirTemp("", "test-src")
		if err != nil {
			t.Fatal(err)
		}
		defer os.RemoveAll(d)
		collector := golangTestCollector{results: []string{}}
		fact, err := GolangTestExistsRule(d, collector, nil)
		if err != nil {
			t.Fatal(err)
		}
		assert.NotEqual(t, fact, GolangTestsExistFact)
	})
}
