package rules

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExpressDependencyExistsRule(t *testing.T) {
	t.Run("ExpressDepedencyExistsFact is true", func(t *testing.T) {
		d, err := os.MkdirTemp("", "test-src")
		if err != nil {
			t.Fatal(err)
		}
		defer os.RemoveAll(d)
		if err := os.WriteFile(filepath.Join(d, "package.json"), []byte(`{"dependencies":{"express": "^4.17.2"}}`), 0744); err != nil {
			t.Fatal(err)
		}
		fact, err := expressDependencyExistsRule(d, nil, nil)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, fact, ExpressDepedencyExistsFact)
	})

	t.Run("ExpressDepedencyExistsFact is false", func(t *testing.T) {
		d, err := os.MkdirTemp("", "test-src")
		if err != nil {
			t.Fatal(err)
		}
		defer os.RemoveAll(d)
		if err := os.WriteFile(filepath.Join(d, "package.json"), []byte(`{"dependencies":{"bcrypt": "^5.0.1"}}`), 0744); err != nil {
			t.Fatal(err)
		}
		fact, err := expressDependencyExistsRule(d, nil, nil)
		if err != nil {
			t.Fatal(err)
		}
		assert.NotEqual(t, fact, ExpressDepedencyExistsFact)
	})
}
