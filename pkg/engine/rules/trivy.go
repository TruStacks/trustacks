package rules

import (
	"os"
	"path/filepath"

	mapset "github.com/deckarep/golang-set/v2"
	"trustacks.io/trustacks/engine"
)

var (
	TrivyConfigExistsFact = engine.NewFact()
)

var trivyConfigExistsRule engine.Rule = func(source string, collector engine.Collector, _ mapset.Set[engine.Fact]) (engine.Fact, error) {
	var fact = engine.NilFact
	if _, err := os.Stat(filepath.Join(source, "trivy.yaml")); os.IsNotExist(err) {
		return fact, nil
	} else if err != nil {
		return fact, err
	}
	fact = TrivyConfigExistsFact
	return fact, nil
}

func init() {
	engine.AddToRuleset(&containerfileHasNoDependenciesRule, &trivyConfigExistsRule)
}
