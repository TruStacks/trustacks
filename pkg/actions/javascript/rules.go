package javascript

import (
	"encoding/json"
	"os"
	"path/filepath"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/trustacks/trustacks/pkg/engine"
)

var (
	// PackageJSONExistsFact is true if the package.json file exists
	// in the root of the application source.
	PackageJSONExistsFact = engine.NewFact()
	// PackageJSONVersionExistsFact is true if the package.json file
	// contains the version key.
	PackageJSONVersionExistsFact = engine.NewFact()
)

// PackageJSONExistsRule checks if the package.json file exits in the
// root of the filesystem.
var PackageJSONExistsRule engine.Rule = func(source string, _ engine.Collector, _ mapset.Set[engine.Fact]) (engine.Fact, error) {
	var fact = engine.NilFact
	if _, err := os.Stat(filepath.Join(source, "package.json")); os.IsNotExist(err) {
		return fact, nil
	} else if err != nil {
		return fact, err
	}
	fact = PackageJSONExistsFact
	return fact, nil
}

// PackageJSONVersionExistsRule checks that the version key exist in
// the package.json configuration file.
var PackageJSONVersionExistsRule engine.Rule = func(source string, _ engine.Collector, _ mapset.Set[engine.Fact]) (engine.Fact, error) {
	var fact = engine.NilFact
	var packageJSON map[string]interface{}
	data, err := os.ReadFile(filepath.Join(source, "package.json"))
	if err != nil {
		return fact, err
	}
	if err := json.Unmarshal(data, &packageJSON); err != nil {
		return fact, err
	}
	if _, ok := packageJSON["version"]; ok {
		fact = PackageJSONVersionExistsFact
	}
	return fact, nil
}

func init() {
	engine.AddToRuleset(&PackageJSONExistsRule, &PackageJSONVersionExistsRule)
}
