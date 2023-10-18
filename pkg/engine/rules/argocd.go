package rules

import (
	"os"
	"regexp"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/trustacks/trustacks/pkg/engine"
	"gopkg.in/yaml.v2"
)

var (
	ArgoCDApplicationExistsFact = engine.NewFact()
)

type ArgoCDApplication struct {
	APIVersion string `yaml:"apiVersion,omitempty"`
	Kind       string `yaml:"kind,omitempty"`
}

var argoCDApplicationExistsRule engine.Rule = func(source string, collector engine.Collector, _ mapset.Set[engine.Fact]) (engine.Fact, error) {
	var fact = engine.NilFact
	matches := collector.Search(".*.yaml")
	for _, path := range matches {
		app := &ArgoCDApplication{}
		contents, err := os.ReadFile(path)
		if err != nil {
			return fact, err
		}
		if err := yaml.Unmarshal(contents, &app); err != nil {
			match, _ := regexp.Match(`: {{ \.Values.*`, contents)
			if !match {
				return fact, err
			}
		}
		if app.APIVersion == "argoproj.io/v1alpha1" && app.Kind == "Application" {
			fact = ArgoCDApplicationExistsFact
		}
	}
	return fact, nil
}

func init() {
	engine.AddToRuleset(&containerfileExistsRule, &argoCDApplicationExistsRule)
}
