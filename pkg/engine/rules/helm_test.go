package rules

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
	"trustacks.io/trustacks/engine"
)

type helmTestCollector struct {
	results []string
}

func (c helmTestCollector) Search(pattern string) []string {
	return c.results
}
func TestHasHelmChartRule(t *testing.T) {
	t.Run("HelmChartExistsFact is true", func(t *testing.T) {
		tests := []struct {
			version string
		}{
			{"v1"},
			{"v2"},
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
			if err := os.MkdirAll(filepath.Join(d, "helm"), 0755); err != nil {
				t.Fatal(err)
			}
			yml, err := yaml.Marshal(map[string]interface{}{"apiVersion": tc.version})
			if err != nil {
				t.Fatal(err)
			}
			if err := os.WriteFile(filepath.Join(d, "helm", "Chart.yaml"), []byte(yml), 0744); err != nil {
				t.Fatal(err)
			}
			collector := &helmTestCollector{results: []string{filepath.Join(d, "helm", "Chart.yaml")}}
			fact, err := helmChartExistsRule(d, collector, nil)
			if err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, fact, HelmChartExistsFact)
		}
	})

	t.Run("HelmChartExistsFact is false", func(t *testing.T) {
		d, err := os.MkdirTemp("", "test-src")
		if err != nil {
			t.Fatal(err)
		}
		fact, err := helmChartExistsRule(d, &helmTestCollector{results: []string{}}, nil)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, fact, engine.NilFact)
		yml, err := yaml.Marshal(map[string]interface{}{"apiVersion": "v255"})
		if err != nil {
			t.Fatal(err)
		}

		if err := os.MkdirAll(filepath.Join(d, "helm"), 0755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(filepath.Join(d, "helm", "Chart.yaml"), []byte(yml), 0744); err != nil {
			t.Fatal(err)
		}
		collector := &helmTestCollector{results: []string{filepath.Join(d, "helm", "Chart.yaml")}}
		fact, err = helmChartExistsRule(d, collector, nil)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, fact, engine.NilFact)
	})
}
