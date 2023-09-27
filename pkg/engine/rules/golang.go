package rules

import (
	"os"
	"path/filepath"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/trustacks/trustacks/pkg/engine"
)

var (
	GolangTestsExistFact = engine.NewFact()
)

var golangTestExistsRule engine.Rule = func(source string, _ engine.Collector, _ mapset.Set[engine.Fact]) (engine.Fact, error) {
	var fact = engine.NilFact
	if _, err := os.Stat(filepath.Join(source, "package.json")); os.IsNotExist(err) {
		return fact, nil
	} else if err != nil {
		return fact, err
	}
	fact = GolangTestsExistFact
	return fact, nil
}
