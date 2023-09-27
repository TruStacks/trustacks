package rules

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/pelletier/go-toml/v2"
	"github.com/trustacks/trustacks/pkg/engine"
)

var (
	PyProjectTomlExistsFact   = engine.NewFact()
	PipRequirementsExistsFact = engine.NewFact()
)

var pyProjectTomlExistsRule engine.Rule = func(source string, _ engine.Collector, _ mapset.Set[engine.Fact]) (engine.Fact, error) {
	var fact = engine.NilFact
	if _, err := os.Stat(filepath.Join(source, "pyproject.toml")); os.IsNotExist(err) {
		return fact, nil
	} else if err != nil {
		return fact, err
	}
	fact = PyProjectTomlExistsFact
	return fact, nil
}

var pipRequirementsExistsRule engine.Rule = func(source string, _ engine.Collector, _ mapset.Set[engine.Fact]) (engine.Fact, error) {
	var fact = engine.NilFact
	if _, err := os.Stat(filepath.Join(source, "requirements.txt")); os.IsNotExist(err) {
		return fact, nil
	} else if err != nil {
		return fact, err
	}
	fact = PipRequirementsExistsFact
	return fact, nil
}

type PoetryTOML struct {
	Packages []PoetryTOMLPackage `toml:"package"`
}

type PoetryTOMLPackage struct {
	Name string `toml:"name"`
}

func checkPoetryDependencies(path, dependency string) (bool, error) {
	var poetry PoetryTOML
	data, err := os.ReadFile(path)
	if err != nil {
		return false, err
	}
	if err := toml.Unmarshal(data, &poetry); err != nil {
		return false, err
	}
	for _, dep := range poetry.Packages {
		if dep.Name == dependency {
			return true, nil
		}
	}
	return false, nil
}

func checkPipRequirements(path, dependency string) (bool, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return false, err
	}
	re, err := regexp.Compile(fmt.Sprintf("^%s==", dependency))
	if err != nil {
		return false, err
	}
	return re.Match(data), nil
}
