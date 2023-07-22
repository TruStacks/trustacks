package engine

import (
	"encoding/json"
	"os"
	"path/filepath"

	mapset "github.com/deckarep/golang-set/v2"
)

type Fact int

const NilFact Fact = -1

const (
	PackageJsonExistsFact Fact = iota
	PackageJsonVersionExistsFact
	ReactScriptsBuildExistsFact
	ReactScriptsTestExistsFact
	EslintConfigExistsFact
	ContainerfileExistFact
	GoModExistsFact
	HasGolangTestSourcesFact
	GolangCILintExistsFact
	GoreleaserExistsFact
	GoreleaserProRequired
)

var packageJsonExistsRule Rule = func(source string, _ *SourceCollector, _ mapset.Set[Fact]) (Fact, error) {
	var fact = NilFact
	if _, err := os.Stat(filepath.Join(source, "package.json")); os.IsNotExist(err) {
		return fact, nil
	} else if err != nil {
		return fact, err
	}
	fact = PackageJsonExistsFact
	return fact, nil
}

var packageJsonVersionExistsRule Rule = func(source string, _ *SourceCollector, _ mapset.Set[Fact]) (Fact, error) {
	var fact = NilFact
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

var reactScriptsTestExsitsRule Rule = func(source string, _ *SourceCollector, _ mapset.Set[Fact]) (Fact, error) {
	var fact = NilFact
	var packageJson map[string]interface{}
	data, err := os.ReadFile(filepath.Join(source, "package.json"))
	if err != nil {
		return fact, err
	}
	if err := json.Unmarshal(data, &packageJson); err != nil {
		return fact, err
	}
	if packageJson["scripts"].(map[string]interface{})["test"] == "react-scripts test" {
		fact = ReactScriptsTestExistsFact
	}
	return fact, nil
}

var reactScriptsBuildExsitsRule Rule = func(source string, _ *SourceCollector, _ mapset.Set[Fact]) (Fact, error) {
	var fact = NilFact
	var packageJson map[string]interface{}
	data, err := os.ReadFile(filepath.Join(source, "package.json"))
	if err != nil {
		return fact, err
	}
	if err := json.Unmarshal(data, &packageJson); err != nil {
		return fact, err
	}
	if packageJson["scripts"].(map[string]interface{})["build"] == "react-scripts build" {
		fact = ReactScriptsBuildExistsFact
	}
	return fact, nil
}

var eslintConfigExistsRule Rule = func(source string, _ *SourceCollector, _ mapset.Set[Fact]) (Fact, error) {
	var fact = NilFact
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
		fact = EslintConfigExistsFact
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
			fact = EslintConfigExistsFact
		}
	}
	return fact, nil
}

var containerfileExistsRule Rule = func(source string, _ *SourceCollector, _ mapset.Set[Fact]) (Fact, error) {
	var fact = NilFact
	for _, file := range []string{"Dockerfile", "Containerfile"} {
		if _, err := os.Stat(filepath.Join(source, file)); !os.IsNotExist(err) {
			fact = ContainerfileExistFact
		}
	}
	return fact, nil
}

var hasGoModRule Rule = func(source string, collector *SourceCollector, _ mapset.Set[Fact]) (Fact, error) {
	var fact = NilFact
	if _, err := os.Stat(filepath.Join(source, "go.mod")); !os.IsNotExist(err) {
		fact = GoModExistsFact
	}
	return fact, nil
}

var hasGolangTestSourcesRule Rule = func(source string, collector *SourceCollector, _ mapset.Set[Fact]) (Fact, error) {
	var fact = NilFact
	matches := collector.Search("_test.go")
	if len(matches) > 0 {
		fact = HasGolangTestSourcesFact
	}
	return fact, nil
}

var golangCILintExistsRule Rule = func(source string, collector *SourceCollector, _ mapset.Set[Fact]) (Fact, error) {
	var fact = NilFact
	var golangciConfig string
	var golangciConfigPatterns = []string{
		".golangci.yml",
		".golangci.yaml",
		".golangci.toml",
		".golangci.json",
	}
	entries, err := os.ReadDir(source)
	if err != nil {
		return fact, err
	}
eslintSourcePatternMatch:
	for _, entry := range entries {
		for _, pattern := range golangciConfigPatterns {
			if pattern == entry.Name() {
				golangciConfig = entry.Name()
				break eslintSourcePatternMatch
			}
		}
	}
	if golangciConfig != "" {
		fact = GolangCILintExistsFact
	}
	return fact, nil
}

var goreleaserExistsRule Rule = func(source string, collector *SourceCollector, _ mapset.Set[Fact]) (Fact, error) {
	var fact = NilFact
	if _, err := os.Stat(filepath.Join(source, ".goreleaser.yaml")); !os.IsNotExist(err) {
		fact = GoreleaserExistsFact
	}
	return fact, nil
}

func init() {
	addToRuleset(&packageJsonExistsRule, &reactScriptsTestExsitsRule)
	addToRuleset(&packageJsonExistsRule, &reactScriptsBuildExsitsRule)
	addToRuleset(&packageJsonExistsRule, &eslintConfigExistsRule)
	addToRuleset(&packageJsonExistsRule, &packageJsonVersionExistsRule)
	addToRuleset(&reactScriptsBuildExsitsRule, &containerfileExistsRule)

	addToRuleset(&hasGoModRule, &hasGolangTestSourcesRule)
	addToRuleset(&hasGoModRule, &golangCILintExistsRule)
	addToRuleset(&hasGoModRule, &goreleaserExistsRule)
}
