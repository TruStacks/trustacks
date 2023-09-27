package rules

import (
	"os"
	"path/filepath"
	"regexp"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/trustacks/trustacks/pkg/engine"
)

var (
	PytestDependencyExistsFact = engine.NewFact()
)

var pytestDependencyExistsRule engine.Rule = func(source string, collector engine.Collector, _ mapset.Set[engine.Fact]) (engine.Fact, error) {
	var err error
	var hasDependency bool
	var hasImport bool
	var hasConfig bool
	var fact = engine.NilFact
	tests := collector.Search("test_.*.py")
	if len(tests) == 0 {
		return fact, nil
	}
	for _, path := range tests {
		content, err := os.ReadFile(path)
		if err != nil {
			return fact, err
		}
		re, err := regexp.Compile(`import pytest`)
		if err != nil {
			return fact, err
		}
		if re.Match(content) {
			hasImport = true
			break
		}
	}
	if _, err := os.Stat(filepath.Join(source, "pytest.ini")); !os.IsNotExist(err) {
		hasConfig = true
	}
	if _, err := os.Stat(filepath.Join(source, "poetry.lock")); !os.IsNotExist(err) {
		hasDependency, err = checkPoetryDependencies(filepath.Join(source, "poetry.lock"), "pytest")
		if err != nil {
			return fact, err
		}
	}
	if _, err := os.Stat(filepath.Join(source, "requirements.txt")); !os.IsNotExist(err) {
		hasDependency, err = checkPipRequirements(filepath.Join(source, "requirements.txt"), "pytest")
		if err != nil {
			return fact, err
		}
	}
	if hasImport || hasConfig || hasDependency {
		fact = PytestDependencyExistsFact
	}
	return fact, err
}

func init() {
	engine.AddToRuleset(&pyProjectTomlExistsRule, &pytestDependencyExistsRule)
	engine.AddToRuleset(&pipRequirementsExistsRule, &pytestDependencyExistsRule)
}
