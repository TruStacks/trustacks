package rules

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/trustacks/trustacks/pkg/engine"
	"gopkg.in/yaml.v2"
)

type helmTestCollector struct {
	results []string
}

func (c helmTestCollector) Search(pattern string) []string {
	return c.results
}
func TestHasHelmChartRule(t *testing.T) {
	t.Run("ArgoCDApplicationExistsFact is true", func(t *testing.T) {
		d, err := os.MkdirTemp("", "test-src")
		if err != nil {
			t.Fatal(err)
		}
		defer os.RemoveAll(d)
		if err := os.MkdirAll(filepath.Join(d, "helm"), 0755); err != nil {
			t.Fatal(err)
		}
		yml, err := yaml.Marshal(map[string]interface{}{"apiVersion": "argoproj.io/v1alpha1", "kind": "Application"})
		if err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(filepath.Join(d, "argo-app.yaml"), []byte(yml), 0744); err != nil {
			t.Fatal(err)
		}
		collector := &helmTestCollector{results: []string{filepath.Join(d, "argo-app.yaml")}}
		fact, err := argoCDApplicationExistsRule(d, collector, nil)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, fact, ArgoCDApplicationExistsFact)
	})

	t.Run("HelmChartExistsFact is false", func(t *testing.T) {
		d, err := os.MkdirTemp("", "test-src")
		if err != nil {
			t.Fatal(err)
		}
		fact, err := argoCDApplicationExistsRule(d, &helmTestCollector{results: []string{}}, nil)
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
		fact, err = argoCDApplicationExistsRule(d, collector, nil)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, fact, engine.NilFact)
	})
}
