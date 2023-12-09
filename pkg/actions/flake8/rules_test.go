package flake8

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFlake8ConfigExistsRule(t *testing.T) {
	t.Run("Flake8ConfigExistsFact is true", func(t *testing.T) {
		tests := []struct {
			configPattern string
		}{
			{".flake8"},
			{"setup.cfg"},
			{"tox.ini"},
		}
		tempDirs := []string{}
		defer func() {
			for _, d := range tempDirs {
				os.RemoveAll(d)
			}
		}()
		for _, tc := range tests {
			d, err := os.MkdirTemp("", "test-src")
			if err != nil {
				t.Fatal(err)
			}
			tempDirs = append(tempDirs, d)
			if err := os.WriteFile(filepath.Join(d, tc.configPattern), []byte("[flake8]"), 0744); err != nil {
				t.Fatal(err)
			}
			fact, err := Flake8ConfigExistsRule(d, nil, nil)
			if err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, fact, Flake8ConfigExistsFact)
		}
	})

	t.Run("Flake8ConfigExistsFact is false", func(t *testing.T) {
		d, err := os.MkdirTemp("", "test-src")
		if err != nil {
			t.Fatal(err)
		}
		defer os.RemoveAll(d)
		if err := os.WriteFile(filepath.Join(d, ".flake8"), []byte("[fake8]"), 0744); err != nil {
			t.Fatal(err)
		}
		fact, err := Flake8ConfigExistsRule(d, nil, nil)
		if err != nil {
			t.Fatal(err)
		}
		assert.NotEqual(t, fact, Flake8ConfigExistsFact)
	})
}
