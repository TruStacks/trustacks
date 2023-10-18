package rules

import (
	"os"
	"path/filepath"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/trustacks/trustacks/pkg/engine"
)

var (
	GoModExistsFact      = engine.NewFact()
	GolangTestsExistFact = engine.NewFact()
)

var goModExistsRule engine.Rule = func(source string, _ engine.Collector, _ mapset.Set[engine.Fact]) (engine.Fact, error) {
	var fact = engine.NilFact
	if _, err := os.Stat(filepath.Join(source, "go.mod")); os.IsNotExist(err) {
		return fact, nil
	} else if err != nil {
		return fact, err
	}
	fact = GoModExistsFact
	return fact, nil
}

var golangTestExistsRule engine.Rule = func(source string, collector engine.Collector, _ mapset.Set[engine.Fact]) (engine.Fact, error) {
	var fact = engine.NilFact
	if len(collector.Search(".*_test.go")) > 0 {
		fact = GolangTestsExistFact
	}
	return fact, nil
}

func init() {
	engine.AddToRuleset(&goModExistsRule, &golangTestExistsRule)
}
