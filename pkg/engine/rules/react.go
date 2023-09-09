package rules

import (
	"encoding/json"
	"os"
	"path/filepath"
	"regexp"

	mapset "github.com/deckarep/golang-set/v2"
	"trustacks.io/trustacks/engine"
)

var (
	// ReactScriptsTestExistsFact is true if react-scripts test is
	// present in the application's package.json.
	ReactScriptsTestExistsFact = engine.NewFact()
	// ReactScriptsBuildExistsFact is true if react-scripts build is
	// present in the application's package.json.
	ReactScriptsBuildExistsFact = engine.NewFact()
)

type reactPackageJsonSpec struct {
	Scripts reactPackageJsonSpecScripts `json:"scripts"`
}

type reactPackageJsonSpecScripts struct {
	Build string `json:"build"`
	Test  string `json:"test"`
}

// reactScriptsTestExsitsRule checks if the react-scripts test
// command is used as the test script in package.json.
var reactScriptsTestExsitsRule engine.Rule = func(source string, collector engine.Collector, _ mapset.Set[engine.Fact]) (engine.Fact, error) {
	var fact = engine.NilFact
	if len(collector.Search(`\.test.js`)) > 0 {
		packageJson := &reactPackageJsonSpec{}
		data, err := os.ReadFile(filepath.Join(source, "package.json"))
		if err != nil {
			return fact, err
		}
		if err := json.Unmarshal(data, packageJson); err != nil {
			return fact, err
		}
		if packageJson.Scripts.Test != "" {
			re, err := regexp.Compile(`(^|\s)react-scripts\s+test`)
			if err != nil {
				return fact, err
			}
			if len(re.FindAllString(packageJson.Scripts.Test, -1)) > 0 {
				return ReactScriptsTestExistsFact, nil
			}
		}
	}
	return fact, nil
}

// reactScriptsBuildExsitsRule checks if the react-scripts build
// command is used as the build script in package.json.
var reactScriptsBuildExsitsRule engine.Rule = func(source string, _ engine.Collector, _ mapset.Set[engine.Fact]) (engine.Fact, error) {
	var fact = engine.NilFact
	packageJson := &reactPackageJsonSpec{}
	data, err := os.ReadFile(filepath.Join(source, "package.json"))
	if err != nil {
		return fact, err
	}
	if err := json.Unmarshal(data, &packageJson); err != nil {
		return fact, err
	}
	if packageJson.Scripts.Build != "" {
		re, err := regexp.Compile(`(^|\s)react-scripts\s+build`)
		if err != nil {
			return fact, err
		}
		if len(re.FindAllString(packageJson.Scripts.Build, -1)) > 0 {
			return ReactScriptsBuildExistsFact, nil
		}
	}
	return fact, nil
}

func init() {
	engine.AddToRuleset(&packageJsonExistsRule, &reactScriptsTestExsitsRule)
	engine.AddToRuleset(&packageJsonExistsRule, &reactScriptsBuildExsitsRule)
}
