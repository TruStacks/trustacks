package rules

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrivyConfigExistsRule(t *testing.T) {
	t.Run("TrivyConfigExistsFact is true", func(t *testing.T) {
		d, err := os.MkdirTemp("", "test-src")
		if err != nil {
			t.Fatal(err)
		}
		defer os.RemoveAll(d)
		if err := os.WriteFile(filepath.Join(d, "trivy.yaml"), []byte(`quiet: true`), 0744); err != nil {
			t.Fatal(err)
		}
		fact, err := trivyConfigExistsRule(d, nil, nil)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, fact, TrivyConfigExistsFact)
	})

	t.Run("TrivyConfigExistsFact is false", func(t *testing.T) {
		d, err := os.MkdirTemp("", "test-src")
		if err != nil {
			t.Fatal(err)
		}
		defer os.RemoveAll(d)
		fact, err := trivyConfigExistsRule(d, nil, nil)
		if err != nil {
			t.Fatal(err)
		}
		assert.NotEqual(t, fact, TrivyConfigExistsFact)
	})
}
