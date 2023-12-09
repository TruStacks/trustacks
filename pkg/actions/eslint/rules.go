package eslint

import (
	"encoding/json"
	"os"
	"path/filepath"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/trustacks/trustacks/pkg/actions/javascript"
	"github.com/trustacks/trustacks/pkg/engine"
)

var (
	ESLintConfigExistsFact = engine.NewFact()
)

var ESLintConfigExistsRule engine.Rule = func(source string, _ engine.Collector, _ mapset.Set[engine.Fact]) (engine.Fact, error) {
	var fact = engine.NilFact
	var eslintConfig string
	var eslintConfigFiles = []string{
		".eslintrc.js",
		".eslintrc.cjs",
		".eslintrc.yaml",
		".eslintrc.yml",
		".eslintrc.json",
		"package.json",
	}
	entries, err := os.ReadDir(source)
	if err != nil {
		return fact, err
	}
	matched := false
	for _, entry := range entries {
		if matched {
			break
		}
		for _, config := range eslintConfigFiles {
			if config == entry.Name() {
				eslintConfig = entry.Name()
				matched = true
			}
		}
	}
	if eslintConfig != "" && eslintConfig != "package.json" {
		fact = ESLintConfigExistsFact
	} else {
		var packageJSON map[string]interface{}
		data, err := os.ReadFile(filepath.Join(source, "package.json"))
		if err != nil {
			return fact, err
		}
		if err := json.Unmarshal(data, &packageJSON); err != nil {
			return fact, err
		}
		if _, ok := packageJSON["eslintConfig"]; ok {
			fact = ESLintConfigExistsFact
		}
	}
	return fact, nil
}

func init() {
	engine.AddToRuleset(&javascript.PackageJSONExistsRule, &ESLintConfigExistsRule)
}
