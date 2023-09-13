package rules

import (
	"encoding/json"
	"os"
	"path/filepath"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/trustacks/trustacks/pkg/engine"
)

var (
	ExpressDepedencyExistsFact = engine.NewFact()
)

type expressPackageJsonSpec struct {
	Dependencies map[string]interface{} `json:"dependencies"`
}

var expressDependencyExistsRule engine.Rule = func(source string, collector engine.Collector, _ mapset.Set[engine.Fact]) (engine.Fact, error) {
	var fact = engine.NilFact
	packageJson := &expressPackageJsonSpec{}
	data, err := os.ReadFile(filepath.Join(source, "package.json"))
	if err != nil {
		return fact, err
	}
	if err := json.Unmarshal(data, &packageJson); err != nil {
		return fact, err
	}
	if _, ok := packageJson.Dependencies["express"]; ok {
		fact = ExpressDepedencyExistsFact
	}
	return fact, nil
}

func init() {
	engine.AddToRuleset(&packageJsonExistsRule, &expressDependencyExistsRule)
}
