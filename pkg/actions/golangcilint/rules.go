package golangcilint

import (
	"fmt"
	"os"
	"path/filepath"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/trustacks/trustacks/pkg/actions/golang"
	"github.com/trustacks/trustacks/pkg/engine"
)

var (
	GolangCILintConfigExistsFact = engine.NewFact()
)

var GolangCILintConfigExistsRule engine.Rule = func(source string, _ engine.Collector, _ mapset.Set[engine.Fact]) (engine.Fact, error) {
	var fact = engine.NilFact
	for _, ext := range []string{"yml", "yaml", "toml", "json"} {
		if _, err := os.Stat(filepath.Join(source, fmt.Sprintf(".golangci.%s", ext))); !os.IsNotExist(err) {
			fact = GolangCILintConfigExistsFact
			break
		}
	}
	return fact, nil
}

func init() {
	engine.AddToRuleset(&golang.GoModExistsRule, &GolangCILintConfigExistsRule)
}
