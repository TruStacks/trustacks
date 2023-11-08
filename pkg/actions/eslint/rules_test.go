package eslint

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestESLintConfigExistsRule(t *testing.T) {
	t.Run("ESLintConfigExistsFact is true", func(t *testing.T) {
		tests := []struct {
			configPattern string
		}{
			{".eslintrc.js"},
			{".eslintrc.cjs"},
			{".eslintrc.yaml"},
			{".eslintrc.yml"},
			{".eslintrc.json"},
			{"package.json"},
		}
		tempDirs := []string{}
		defer func() {
			for _, d := range tempDirs {
				os.RemoveAll(d)
			}
		}()
		for _, tc := range tests {
			config := "{}"
			// add the eslint config key for package.json
			if tc.configPattern == "package.json" {
				config = `{"eslintConfig":{}}`
			}
			d, err := os.MkdirTemp("", "test-src")
			if err != nil {
				t.Fatal(err)
			}
			tempDirs = append(tempDirs, d)
			if err := os.WriteFile(filepath.Join(d, tc.configPattern), []byte(config), 0744); err != nil {
				t.Fatal(err)
			}
			fact, err := ESLintConfigExistsRule(d, nil, nil)
			if err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, fact, ESLintConfigExistsFact)
		}
	})

	t.Run("ESLintConfigExistsFact is false", func(t *testing.T) {
		d, err := os.MkdirTemp("", "test-src")
		if err != nil {
			t.Fatal(err)
		}
		defer os.RemoveAll(d)
		if err := os.WriteFile(filepath.Join(d, "package.json"), []byte("{}"), 0744); err != nil {
			t.Fatal(err)
		}
		fact, err := ESLintConfigExistsRule(d, nil, nil)
		if err != nil {
			t.Fatal(err)
		}
		assert.NotEqual(t, fact, ESLintConfigExistsFact)
	})
}
