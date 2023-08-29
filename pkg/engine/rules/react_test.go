package rules

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

type reactScriptsTestCollector struct {
	results []string
}

func (c reactScriptsTestCollector) Search(pattern string) []string {
	return c.results
}

func TestReactScriptsTestExsitsRule(t *testing.T) {
	t.Run("ReactScriptsTestExistsFact is true", func(t *testing.T) {
		tests := []struct {
			Command string
		}{
			{"react-scripts test"},
			{"react-scripts                       test                   "},
			{"react-scripts test --ci --watchAll=false"},
			{"ENV_VAR=test react-scripts test"},
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
			if err := os.WriteFile(filepath.Join(d, "package.json"), []byte(fmt.Sprintf(`{"scripts":{"test": "%s"}}`, tc.Command)), 0744); err != nil {
				t.Fatal(err)
			}
			collector := reactScriptsTestCollector{results: []string{"my.test.js"}}
			fact, err := reactScriptsTestExsitsRule(d, collector, nil)
			if err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, fact, ReactScriptsTestExistsFact)
		}
	})

	t.Run("ReactScriptsTestExistsFact is false", func(t *testing.T) {
		tests := []struct {
			Command string
		}{
			{"react -scripts test"},
			{"ENV=react-scripts test --ci --watchAll=false"},
			{"react-scripts --ci test --watchAll=false"},
			{"ENV_VAR=test test react-scripts"},
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
			if err := os.WriteFile(filepath.Join(d, "package.json"), []byte(fmt.Sprintf(`{"scripts":{"test": "%s"}}`, tc.Command)), 0744); err != nil {
				t.Fatal(err)
			}
			collector := reactScriptsTestCollector{results: []string{"my.test.js"}}
			fact, err := reactScriptsTestExsitsRule(d, collector, nil)
			if err != nil {
				t.Fatal(err)
			}
			assert.NotEqual(t, fact, ReactScriptsTestExistsFact)
		}
	})
}

func TestReactScriptsBuildExsitsRule(t *testing.T) {
	t.Run("ReactScriptsBuildExistsFact is true", func(t *testing.T) {
		tests := []struct {
			Command string
		}{
			{"react-scripts build"},
			{"react-scripts                       build                   "},
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
			if err := os.WriteFile(filepath.Join(d, "package.json"), []byte(fmt.Sprintf(`{"scripts":{"build": "%s"}}`, tc.Command)), 0744); err != nil {
				t.Fatal(err)
			}
			fact, err := reactScriptsBuildExsitsRule(d, nil, nil)
			if err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, fact, ReactScriptsBuildExistsFact)
		}
	})

	t.Run("ReactScriptsBuildExistsFact is false", func(t *testing.T) {
		tests := []struct {
			Command string
		}{
			{"react -scripts build"},
			{"ENV_VAR=react-scripts build"},
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
			if err := os.WriteFile(filepath.Join(d, "package.json"), []byte(fmt.Sprintf(`{"scripts":{"build": "%s"}}`, tc.Command)), 0744); err != nil {
				t.Fatal(err)
			}
			fact, err := reactScriptsBuildExsitsRule(d, nil, nil)
			if err != nil {
				t.Fatal(err)
			}
			assert.NotEqual(t, fact, ReactScriptsBuildExistsFact)
		}
	})
}
