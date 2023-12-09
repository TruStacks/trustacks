package container

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/trustacks/trustacks/pkg/engine"
)

var (
	ContainerfileExistFact                      = engine.NewFact()
	ContainerfileHasPredictableDependenciesFact = engine.NewFact()
)

var ContainerfileExistsRule engine.Rule = func(source string, _ engine.Collector, _ mapset.Set[engine.Fact]) (engine.Fact, error) {
	var fact = engine.NilFact
	for _, file := range []string{"Dockerfile", "Containerfile"} {
		if _, err := os.Stat(filepath.Join(source, file)); !os.IsNotExist(err) {
			fact = ContainerfileExistFact
		}
	}
	return fact, nil
}

var ContainerfileHasNoDependenciesRule engine.Rule = func(source string, _ engine.Collector, _ mapset.Set[engine.Fact]) (engine.Fact, error) {
	var fact = engine.NilFact
	for _, file := range []string{"Dockerfile", "Containerfile"} {
		if _, err := os.Stat(filepath.Join(source, file)); !os.IsNotExist(err) {
			contents, err := os.ReadFile(filepath.Join(source, file))
			if err != nil {
				return fact, err
			}
			re := regexp.MustCompile(`COPY\s(.*?)\s`)
			matches := re.FindAllStringSubmatch(string(contents), -1)
			for _, match := range matches {
				copyCmd := match[1]
				if copyCmd == "." || strings.Contains(copyCmd, "--from=") {
					continue
				}
				if _, err := os.Stat(filepath.Join(source, copyCmd)); os.IsNotExist(err) {
					return fact, nil
				}
			}
		}
	}
	fact = ContainerfileHasPredictableDependenciesFact
	return fact, nil
}

var ContainerfileHasBuildCopyRule engine.Rule = func(source string, _ engine.Collector, _ mapset.Set[engine.Fact]) (engine.Fact, error) {
	var fact = engine.NilFact
	if _, err := os.Stat(filepath.Join(source, ".build")); !os.IsNotExist(err) {
		return fact, nil
	}
	for _, file := range []string{"Dockerfile", "Containerfile"} {
		if _, err := os.Stat(filepath.Join(source, file)); !os.IsNotExist(err) {
			contents, err := os.ReadFile(filepath.Join(source, file))
			if err != nil {
				return fact, err
			}
			re := regexp.MustCompile(`COPY\s*.build\s*`)
			if !re.Match(contents) {
				return fact, err
			}
		}
	}
	fact = ContainerfileHasPredictableDependenciesFact
	return fact, nil
}

func init() {
	engine.AddToRuleset(&ContainerfileExistsRule, &ContainerfileHasNoDependenciesRule)
	engine.AddToRuleset(&ContainerfileExistsRule, &ContainerfileHasBuildCopyRule)
}
