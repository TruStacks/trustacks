package rules

import (
	"os"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/trustacks/pkg/engine"
	"gopkg.in/yaml.v2"
)

var (
	HelmChartExistsFact = engine.NewFact()
)

var helmChartExistsRule engine.Rule = func(source string, collector engine.Collector, _ mapset.Set[engine.Fact]) (engine.Fact, error) {
	var fact = engine.NilFact
	matches := collector.Search("Chart.yaml")
	for _, path := range matches {
		chart := map[string]interface{}{}
		contents, err := os.ReadFile(path)
		if err != nil {
			return fact, err
		}
		if err := yaml.Unmarshal(contents, &chart); err != nil {
			return fact, err
		}
		if apiVersion, ok := chart["apiVersion"]; ok {
			if apiVersion != "v1" && apiVersion != "v2" {
				return fact, err
			}
		}
		fact = HelmChartExistsFact
	}
	return fact, nil
}

func init() {
	engine.AddToRuleset(&containerfileExistsRule, &helmChartExistsRule)
}
