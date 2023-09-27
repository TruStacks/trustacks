package rules

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGolangCILintConfigExistsRule(t *testing.T) {
	t.Run("GolangCILintConfigExistsFact is true", func(t *testing.T) {
		tempDirs := []string{}
		defer func() {
			for _, d := range tempDirs {
				os.RemoveAll(d)
			}
		}()
		for _, ext := range []string{"yml", "yaml", "toml", "json"} {
			d, err := os.MkdirTemp("", "test-src")
			if err != nil {
				t.Fatal(err)
			}
			tempDirs = append(tempDirs, d)
			if err := os.WriteFile(filepath.Join(d, fmt.Sprintf(".golangci.%s", ext)), []byte(""), 0744); err != nil {
				t.Fatal(err)
			}
			fact, err := golangCILintConfigExistsRule(d, nil, nil)
			if err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, fact, GolangCILintConfigExistsFact)
		}
	})

	t.Run("GolangCILintConfigExistsFact is false", func(t *testing.T) {
		d, err := os.MkdirTemp("", "test-src")
		if err != nil {
			t.Fatal(err)
		}
		defer os.RemoveAll(d)
		fact, err := golangCILintConfigExistsRule(d, nil, nil)
		if err != nil {
			t.Fatal(err)
		}
		assert.NotEqual(t, fact, GolangCILintConfigExistsFact)
	})
}
