package npm

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/trustacks/trustacks/pkg/actions/javascript"
	"github.com/trustacks/trustacks/pkg/engine"
)

var (
	NpmTestExistsFact  = engine.NewFact()
	NpmBuildExistsFact = engine.NewFact()
)

type npmPackageJSONSpec struct {
	Scripts npmPackageJSONSpecScripts `json:"scripts"`
}

type npmPackageJSONSpecScripts struct {
	Build *string `json:"build,omitempty"`
	Test  *string `json:"test,omitempty"`
}

var NpmTestExsitsRule engine.Rule = func(source string, collector engine.Collector, _ mapset.Set[engine.Fact]) (engine.Fact, error) {
	var fact = engine.NilFact
	for _, ext := range []string{"js", "jsx", "ts", "tsx"} {
		if len(collector.Search(fmt.Sprintf(".*.test.%s", ext))) > 0 {
			packageJSON := &npmPackageJSONSpec{}
			data, err := os.ReadFile(filepath.Join(source, "package.json"))
			if err != nil {
				return fact, err
			}
			if err := json.Unmarshal(data, packageJSON); err != nil {
				return fact, err
			}
			if packageJSON.Scripts.Test != nil {
				fact = NpmTestExistsFact
			}
			break
		}
	}
	return fact, nil
}

var NpmBuildExsitsRule engine.Rule = func(source string, _ engine.Collector, _ mapset.Set[engine.Fact]) (engine.Fact, error) {
	var fact = engine.NilFact
	packageJSON := &npmPackageJSONSpec{}
	data, err := os.ReadFile(filepath.Join(source, "package.json"))
	if err != nil {
		return fact, err
	}
	if err := json.Unmarshal(data, &packageJSON); err != nil {
		return fact, err
	}
	if packageJSON.Scripts.Build != nil {
		fact = NpmBuildExistsFact
	}
	return fact, nil
}

func init() {
	engine.AddToRuleset(&javascript.PackageJSONExistsRule, &NpmTestExsitsRule)
	engine.AddToRuleset(&javascript.PackageJSONExistsRule, &NpmBuildExsitsRule)
}
