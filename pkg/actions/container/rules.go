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
	ContainerfileExistFact             = engine.NewFact()
	ContainerfileHasNoDependenciesFact = engine.NewFact()
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

func init() {
	engine.AddToRuleset(&ContainerfileExistsRule, &ContainerfileHasNoDependenciesRule)
}
