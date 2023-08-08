package engine

import (
	"encoding/json"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	mapset "github.com/deckarep/golang-set/v2"
	"gopkg.in/yaml.v2"
)

type Fact int

const NilFact Fact = -1

var factInc = 0

func NewFact() Fact {
	factInc++
	return Fact(factInc)
}

var (
	PackageJsonExistsFact              = NewFact()
	PackageJsonVersionExistsFact       = NewFact()
	ReactScriptsBuildExistsFact        = NewFact()
	ReactScriptsTestExistsFact         = NewFact()
	EslintConfigExistsFact             = NewFact()
	ContainerfileExistFact             = NewFact()
	ContainerfileHasNoDependenciesFact = NewFact()
	GoModExistsFact                    = NewFact()
	HasGolangTestSourcesFact           = NewFact()
	GolangCILintExistsFact             = NewFact()
	GoreleaserExistsFact               = NewFact()
	ExpressDepedencyExistsFact         = NewFact()
	HasHelmChartFact                   = NewFact()
	TrivyConfigExistsFact              = NewFact()
	SonarProjectPropertiesExists       = NewFact()
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

var reactScriptsTestExsitsRule Rule = func(source string, collector *SourceCollector, _ mapset.Set[Fact]) (Fact, error) {
	var fact = NilFact
	var packageJson map[string]interface{}
	data, err := os.ReadFile(filepath.Join(source, "package.json"))
	if err != nil {
		return fact, err
	}
	if err := json.Unmarshal(data, &packageJson); err != nil {
		return fact, err
	}
	if len(collector.Search(`\.test.js`)) > 0 {
		if _, ok := packageJson["scripts"].(map[string]interface{}); ok {
			if _, ok := packageJson["scripts"].(map[string]interface{})["test"]; ok {
				if strings.Contains(packageJson["scripts"].(map[string]interface{})["test"].(string), "react-scripts test") {
					fact = ReactScriptsTestExistsFact
				}
			}
		}
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
	if scripts, ok := packageJson["scripts"]; ok {
		if build, ok := scripts.(map[string]interface{})["build"]; ok {
			buildString, ok := build.(string)
			if ok && strings.Contains(buildString, "react-scripts") && strings.Contains(buildString, "build") {
				fact = ReactScriptsBuildExistsFact
			}
		}
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

var ContainerfileHasNoDependenciesRule Rule = func(source string, _ *SourceCollector, _ mapset.Set[Fact]) (Fact, error) {
	var fact = NilFact
	for _, file := range []string{"Dockerfile", "Containerfile"} {
		if _, err := os.Stat(filepath.Join(source, file)); !os.IsNotExist(err) {
			contents, err := os.ReadFile(filepath.Join(source, file))
			if err != nil {
				return fact, err
			}
			re, err := regexp.Compile(`COPY\s(.*?)\s`)
			if err != nil {
				return fact, err
			}
			matches := re.FindAllStringSubmatch(string(contents), -1)
			for _, match := range matches {
				copy := string(match[1])
				if copy == "." || strings.Contains(copy, "--from=") {
					continue
				}
				if _, err := os.Stat(filepath.Join(source, copy)); os.IsNotExist(err) {
					return fact, nil
				}
			}
		}
	}
	fact = ContainerfileHasNoDependenciesFact
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

var expressDepedencyExistsRule Rule = func(source string, collector *SourceCollector, _ mapset.Set[Fact]) (Fact, error) {
	var fact = NilFact
	var packageJson map[string]interface{}
	data, err := os.ReadFile(filepath.Join(source, "package.json"))
	if err != nil {
		return fact, err
	}
	if err := json.Unmarshal(data, &packageJson); err != nil {
		return fact, err
	}
	if deps, ok := packageJson["dependencies"]; ok {
		if _, ok := deps.(map[string]interface{})["express"]; ok {
			fact = ExpressDepedencyExistsFact
		}
	}
	return fact, nil
}

var hasHelmChartRule Rule = func(source string, collector *SourceCollector, _ mapset.Set[Fact]) (Fact, error) {
	var fact = NilFact
	matches := collector.Search("Chart.yaml")
	for _, path := range matches {
		chart := map[string]interface{}{}
		contents, err := os.ReadFile(path)
		if err != nil {
			return fact, err
		}
		if err := yaml.Unmarshal(contents, &chart); err != nil {
			return fact, err
		}
		if apiVersion, ok := chart["apiVersion"]; ok {
			if apiVersion != "v1" && apiVersion != "v2" {
				return fact, err
			}
		}
		fact = HasHelmChartFact
	}
	return fact, nil
}

var trivyConfigExistsRule Rule = func(source string, collector *SourceCollector, _ mapset.Set[Fact]) (Fact, error) {
	var fact = NilFact
	if _, err := os.Stat(filepath.Join(source, "trivy.yaml")); os.IsNotExist(err) {
		return fact, nil
	} else if err != nil {
		return fact, err
	}
	fact = TrivyConfigExistsFact
	return fact, nil
}

var sonarProjectPropertiesExistsRule Rule = func(source string, collector *SourceCollector, _ mapset.Set[Fact]) (Fact, error) {
	var fact = NilFact
	if _, err := os.Stat(filepath.Join(source, "sonar-project.properties")); os.IsNotExist(err) {
		return fact, nil
	} else if err != nil {
		return fact, err
	}
	fact = SonarProjectPropertiesExists
	return fact, nil
}

func init() {
	// packageJsonExistsRule branch.
	addToRuleset(&packageJsonExistsRule, &reactScriptsTestExsitsRule)
	addToRuleset(&packageJsonExistsRule, &reactScriptsBuildExsitsRule)
	addToRuleset(&packageJsonExistsRule, &eslintConfigExistsRule)
	addToRuleset(&packageJsonExistsRule, &packageJsonVersionExistsRule)
	addToRuleset(&packageJsonExistsRule, &expressDepedencyExistsRule)

	// expressDepedencyExistsRule branch.
	addToRuleset(&expressDepedencyExistsRule, &containerfileExistsRule)

	// reactScriptsBuildExsitsRule branch.
	addToRuleset(&reactScriptsBuildExsitsRule, &containerfileExistsRule)

	// containerfileExistsRule branch.
	addToRuleset(&containerfileExistsRule, &ContainerfileHasNoDependenciesRule)
	addToRuleset(&containerfileExistsRule, &hasHelmChartRule)

	//ContainerfileHasNoDependenciesRule branch.
	addToRuleset(&ContainerfileHasNoDependenciesRule, &trivyConfigExistsRule)

	// hasGoModRule branch.
	addToRuleset(&hasGoModRule, &hasGolangTestSourcesRule)
	addToRuleset(&hasGoModRule, &golangCILintExistsRule)
	addToRuleset(&hasGoModRule, &goreleaserExistsRule)

	// sonarProjectPropertiesExistsRule branch.
	addToRuleset(&sonarProjectPropertiesExistsRule, nil)
}
