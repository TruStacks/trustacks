package goreleaser

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGoreleaserConfigExistsRule(t *testing.T) {
	t.Run("GoreleaserConfigExistsFact is true", func(t *testing.T) {
		d, err := os.MkdirTemp("", "test-src")
		if err != nil {
			t.Fatal(err)
		}
		defer os.RemoveAll(d)
		if err := os.WriteFile(filepath.Join(d, ".goreleaser.yaml"), []byte(""), 0744); err != nil {
			t.Fatal(err)
		}
		fact, err := GoreleaserConfigExistsRule(d, nil, nil)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, fact, GoreleaserConfigExistsFact)
	})

	t.Run("GoreleaserConfigExistsFact is false", func(t *testing.T) {
		d, err := os.MkdirTemp("", "test-src")
		if err != nil {
			t.Fatal(err)
		}
		defer os.RemoveAll(d)
		fact, err := GoreleaserConfigExistsRule(d, nil, nil)
		if err != nil {
			t.Fatal(err)
		}
		assert.NotEqual(t, fact, GoreleaserConfigExistsFact)
	})
}
