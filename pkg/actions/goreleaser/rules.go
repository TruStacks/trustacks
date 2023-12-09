package goreleaser

import (
	"os"
	"path/filepath"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/trustacks/trustacks/pkg/actions/golang"
	"github.com/trustacks/trustacks/pkg/engine"
)

var (
	GoreleaserConfigExistsFact = engine.NewFact()
)

var GoreleaserConfigExistsRule engine.Rule = func(source string, _ engine.Collector, _ mapset.Set[engine.Fact]) (engine.Fact, error) {
	var fact = engine.NilFact
	if _, err := os.Stat(filepath.Join(source, ".goreleaser.yaml")); !os.IsNotExist(err) {
		fact = GoreleaserConfigExistsFact
	}
	return fact, nil
}

func init() {
	engine.AddToRuleset(&golang.GoModExistsRule, &GoreleaserConfigExistsRule)
}
