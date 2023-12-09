package trivy

import (
	"os"
	"path/filepath"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/trustacks/trustacks/pkg/actions/container"
	"github.com/trustacks/trustacks/pkg/engine"
)

var (
	TrivyConfigExistsFact = engine.NewFact()
)

var TrivyConfigExistsRule engine.Rule = func(source string, collector engine.Collector, _ mapset.Set[engine.Fact]) (engine.Fact, error) {
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
	engine.AddToRuleset(&container.ContainerfileHasNoDependenciesRule, &TrivyConfigExistsRule)
	engine.AddToRuleset(&container.ContainerfileHasBuildCopyRule, &TrivyConfigExistsRule)
}
