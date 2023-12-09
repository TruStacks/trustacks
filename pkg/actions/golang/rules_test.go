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

func (c golangTestCollector) Search(_ string) []string {
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
		assert.Equal(t, fact, GolangTestsExistsFact)
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
		assert.NotEqual(t, fact, GolangTestsExistsFact)
	})
}

func TestGolangIntegrationTestsExistRule(t *testing.T) {
	t.Run("GolangIntegrationTestsExistFact is true", func(t *testing.T) {
		d, err := os.MkdirTemp("", "test-src")
		if err != nil {
			t.Fatal(err)
		}
		defer os.RemoveAll(d)
		if err := os.WriteFile(
			filepath.Join(d, "my_test.go"),
			[]byte("func TestSomethingIntegration(t *testing.T) {}"),
			0744,
		); err != nil {
			t.Fatal(err)
		}
		collector := golangTestCollector{results: []string{filepath.Join(d, "my_test.go")}}
		fact, err := GolangIntegrationTestExistsRule(d, collector, nil)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, fact, GolangIntegrationTestsExistsFact)
	})

	t.Run("GolangIntegrationTestsExistFact is false", func(t *testing.T) {
		d, err := os.MkdirTemp("", "test-src")
		if err != nil {
			t.Fatal(err)
		}
		defer os.RemoveAll(d)
		if err := os.WriteFile(
			filepath.Join(d, "my_test.go"),
			[]byte("func TestIntegrationSomething(t *testing.T) {}"),
			0744,
		); err != nil {
			t.Fatal(err)
		}
		collector := golangTestCollector{results: []string{filepath.Join(d, "my_test.go")}}
		fact, err := GolangIntegrationTestExistsRule(d, collector, nil)
		if err != nil {
			t.Fatal(err)
		}
		assert.NotEqual(t, fact, GolangIntegrationTestsExistsFact)
	})
}

func TestGolangCmdExistsRule(t *testing.T) {
	t.Run("GolangCmdExistsFact is true", func(t *testing.T) {
		d, err := os.MkdirTemp("", "test-src")
		if err != nil {
			t.Fatal(err)
		}
		defer os.RemoveAll(d)
		if err := os.Mkdir(filepath.Join(d, "cmd"), 0755); err != nil {
			t.Fatal(err)
		}
		fact, err := GolangCmdExistsRule(d, nil, nil)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, fact, GolangCmdExistsFact)
	})

	t.Run("GolangCmdExistsFact is false if cmd is not a directory", func(t *testing.T) {
		d, err := os.MkdirTemp("", "test-src")
		if err != nil {
			t.Fatal(err)
		}
		defer os.RemoveAll(d)
		if err := os.WriteFile(filepath.Join(d, "cmd"), []byte(""), 0644); err != nil {
			t.Fatal(err)
		}
		fact, err := GolangCmdExistsRule(d, nil, nil)
		if err != nil {
			t.Fatal(err)
		}
		assert.NotEqual(t, fact, GolangCmdExistsFact)
	})

	t.Run("GolangCmdExistsFact is false if cmd does not exist", func(t *testing.T) {
		d, err := os.MkdirTemp("", "test-src")
		if err != nil {
			t.Fatal(err)
		}
		defer os.RemoveAll(d)
		fact, err := GolangCmdExistsRule(d, nil, nil)
		if err != nil {
			t.Fatal(err)
		}
		assert.NotEqual(t, fact, GolangCmdExistsFact)
	})
}
