package rules

import (
	"encoding/json"
	"os"
	"path/filepath"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/trustacks/trustacks/pkg/engine"
)

var (
	ESLintConfigExistsFact = engine.NewFact()
)

var eslintConfigExistsRule engine.Rule = func(source string, _ engine.Collector, _ mapset.Set[engine.Fact]) (engine.Fact, error) {
	var fact = engine.NilFact
	var eslintConfig string
	var eslintConfigPatterns = []string{
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
eslintSourcePatternMatch:
	for _, entry := range entries {
		for _, pattern := range eslintConfigPatterns {
			if pattern == entry.Name() {
				eslintConfig = entry.Name()
				break eslintSourcePatternMatch
			}
		}
	}
	if eslintConfig != "" && eslintConfig != "package.json" {
		fact = ESLintConfigExistsFact
	} else {
		var packageJson map[string]interface{}
		data, err := os.ReadFile(filepath.Join(source, "package.json"))
		if err != nil {
			return fact, err
		}
		if err := json.Unmarshal(data, &packageJson); err != nil {
			return fact, err
		}
		if _, ok := packageJson["eslintConfig"]; ok {
			fact = ESLintConfigExistsFact
		}
	}
	return fact, nil
}

func init() {
	engine.AddToRuleset(&expressDependencyExistsRule, &eslintConfigExistsRule)
	engine.AddToRuleset(&reactScriptsBuildExsitsRule, &eslintConfigExistsRule)
}
