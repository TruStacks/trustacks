package rules

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

type pytestTestCollector struct {
	pattern string
	results []string
}

func (c *pytestTestCollector) Search(pattern string) []string {
	c.pattern = pattern
	return c.results
}

func TestPytestDependencyExistsRule(t *testing.T) {
	t.Run("PytestDependencyExistsFact is true", func(t *testing.T) {
		d, err := os.MkdirTemp("", "test-src")
		if err != nil {
			t.Fatal(err)
		}
		defer os.RemoveAll(d)
		if err := os.WriteFile(filepath.Join(d, "test_my.py"), []byte("import pytest"), 0744); err != nil {
			t.Fatal(err)
		}
		collector := &pytestTestCollector{results: []string{filepath.Join(d, "test_my.py")}}
		fact, err := pytestDependencyExistsRule(d, collector, nil)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, "test_.*.py", collector.pattern)
		assert.Equal(t, fact, PytestDependencyExistsFact)
	})

	t.Run("PytestDependencyExistsFact is false", func(t *testing.T) {
		d, err := os.MkdirTemp("", "test-src")
		if err != nil {
			t.Fatal(err)
		}
		defer os.RemoveAll(d)
		if err := os.WriteFile(filepath.Join(d, "test_my.py"), []byte("import unittest"), 0744); err != nil {
			t.Fatal(err)
		}
		collector := &pytestTestCollector{results: []string{filepath.Join(d, "test_my.py")}}
		fact, err := pytestDependencyExistsRule(d, collector, nil)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, "test_.*.py", collector.pattern)
		assert.NotEqual(t, fact, PytestDependencyExistsFact)
	})
}
