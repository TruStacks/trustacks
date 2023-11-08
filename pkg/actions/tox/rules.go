package tox

import (
	"os"
	"path/filepath"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/trustacks/trustacks/pkg/actions/python"
	"github.com/trustacks/trustacks/pkg/engine"
)

var (
	ToxIniExistsFact = engine.NewFact()
)

var ToxIniExistsRule engine.Rule = func(source string, _ engine.Collector, _ mapset.Set[engine.Fact]) (engine.Fact, error) {
	var fact = engine.NilFact
	if _, err := os.Stat(filepath.Join(source, "tox.ini")); os.IsNotExist(err) {
		return fact, nil
	} else if err != nil {
		return fact, err
	}
	fact = ToxIniExistsFact
	return fact, nil
}

func init() {
	engine.AddToRuleset(&python.PyProjectTomlExistsRule, &ToxIniExistsRule)
	engine.AddToRuleset(&python.PipRequirementsExistsRule, &ToxIniExistsRule)
}
