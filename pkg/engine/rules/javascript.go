package rules

import (
	"encoding/json"
	"os"
	"path/filepath"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/trustacks/pkg/engine"
)

var (
	// PackageJsonExistsFact is true if the package.json file exists
	// in the root of the application source.
	PackageJsonExistsFact = engine.NewFact()
	// PackageJsonVersionExistsFact is true if the package.json file
	// contains the version key.
	PackageJsonVersionExistsFact = engine.NewFact()
)

// packageJsonExistsRule checks if the package.json file exits in the
// root of the filesystem.
var packageJsonExistsRule engine.Rule = func(source string, _ engine.Collector, _ mapset.Set[engine.Fact]) (engine.Fact, error) {
	var fact = engine.NilFact
	if _, err := os.Stat(filepath.Join(source, "package.json")); os.IsNotExist(err) {
		return fact, nil
	} else if err != nil {
		return fact, err
	}
	fact = PackageJsonExistsFact
	return fact, nil
}

// packageJsonVersionExistsRule checks that the version key exist in
// the package.json configuration file.
var packageJsonVersionExistsRule engine.Rule = func(source string, _ engine.Collector, _ mapset.Set[engine.Fact]) (engine.Fact, error) {
	var fact = engine.NilFact
	var packageJson map[string]interface{}
	data, err := os.ReadFile(filepath.Join(source, "package.json"))
	if err != nil {
		return fact, err
	}
	if err := json.Unmarshal(data, &packageJson); err != nil {
		return fact, err
	}
	if _, ok := packageJson["version"]; ok {
		fact = PackageJsonVersionExistsFact
	}
	return fact, nil
}

func init() {
	engine.AddToRuleset(&packageJsonExistsRule, &packageJsonVersionExistsRule)
}
