package argocd

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/trustacks/trustacks/pkg/engine"
	"gopkg.in/yaml.v2"
)

type argocdTestCollector struct {
	results []string
}

func (c argocdTestCollector) Search(_ string) []string {
	return c.results
}
func TestArgoCDApplicationExistsRule(t *testing.T) {
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
		if err := os.WriteFile(filepath.Join(d, "argo-app.yaml"), yml, 0744); err != nil {
			t.Fatal(err)
		}
		collector := &argocdTestCollector{results: []string{filepath.Join(d, "argo-app.yaml")}}
		fact, err := ArgoCDApplicationExistsRule(d, collector, nil)
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
		fact, err := ArgoCDApplicationExistsRule(d, &argocdTestCollector{results: []string{}}, nil)
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
		if err := os.WriteFile(filepath.Join(d, "helm", "Chart.yaml"), yml, 0744); err != nil {
			t.Fatal(err)
		}
		collector := &argocdTestCollector{results: []string{filepath.Join(d, "helm", "Chart.yaml")}}
		fact, err = ArgoCDApplicationExistsRule(d, collector, nil)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, fact, engine.NilFact)
	})
}
