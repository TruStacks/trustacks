package rules

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToxIniExistsRule(t *testing.T) {
	t.Run("ToxIniExistsFact is true", func(t *testing.T) {
		d, err := os.MkdirTemp("", "test-src")
		if err != nil {
			t.Fatal(err)
		}
		defer os.RemoveAll(d)
		if err := os.WriteFile(filepath.Join(d, "tox.ini"), []byte(``), 0744); err != nil {
			t.Fatal(err)
		}
		fact, err := toxIniExistsRule(d, nil, nil)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, fact, ToxIniExistsFact)
	})

	t.Run("ToxIniExistsFact is false", func(t *testing.T) {
		d, err := os.MkdirTemp("", "test-src")
		if err != nil {
			t.Fatal(err)
		}
		defer os.RemoveAll(d)
		fact, err := toxIniExistsRule(d, nil, nil)
		if err != nil {
			t.Fatal(err)
		}
		assert.NotEqual(t, fact, ToxIniExistsFact)
	})
}
