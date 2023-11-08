package argocd

import (
	"os"
	"regexp"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/trustacks/trustacks/pkg/actions/container"
	"github.com/trustacks/trustacks/pkg/engine"
	"gopkg.in/yaml.v2"
)

var (
	ArgoCDApplicationExistsFact = engine.NewFact()
)

type ArgoCDApplicationSpec struct {
	APIVersion string                        `yaml:"apiVersion,omitempty"`
	Kind       string                        `yaml:"kind,omitempty"`
	Metadata   ArgoCDApplicationSpecMetadata `yaml:"metadata"`
}

type ArgoCDApplicationSpecMetadata struct {
	Name string `yaml:"name"`
}

var ArgoCDApplicationExistsRule engine.Rule = func(source string, collector engine.Collector, _ mapset.Set[engine.Fact]) (engine.Fact, error) {
	var fact = engine.NilFact
	matches := collector.Search(".*.yaml")
	for _, path := range matches {
		app := &ArgoCDApplicationSpec{}
		contents, err := os.ReadFile(path)
		if err != nil {
			return fact, err
		}
		re, err := regexp.Compile(`: {{ \.Values.*`)
		if err != nil {
			return fact, err
		}
		if err := yaml.Unmarshal(contents, &app); err != nil {
			if !re.Match(contents) {
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
	engine.AddToRuleset(&container.ContainerfileExistsRule, &ArgoCDApplicationExistsRule)
}
