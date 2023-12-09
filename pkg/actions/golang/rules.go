package golang

import (
	"os"
	"path/filepath"
	"regexp"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/trustacks/trustacks/pkg/engine"
)

var (
	GoModExistsFact                  = engine.NewFact()
	GolangTestsExistsFact            = engine.NewFact()
	GolangIntegrationTestsExistsFact = engine.NewFact()
	GolangCmdExistsFact              = engine.NewFact()
)

var GoModExistsRule engine.Rule = func(source string, _ engine.Collector, _ mapset.Set[engine.Fact]) (engine.Fact, error) {
	var fact = engine.NilFact
	if _, err := os.Stat(filepath.Join(source, "go.mod")); os.IsNotExist(err) {
		return fact, nil
	} else if err != nil {
		return fact, err
	}
	fact = GoModExistsFact
	return fact, nil
}

var GolangTestExistsRule engine.Rule = func(source string, collector engine.Collector, _ mapset.Set[engine.Fact]) (engine.Fact, error) {
	var fact = engine.NilFact
	if len(collector.Search(".*_test.go")) > 0 {
		fact = GolangTestsExistsFact
	}
	return fact, nil
}

var GolangIntegrationTestExistsRule engine.Rule = func(source string, collector engine.Collector, _ mapset.Set[engine.Fact]) (engine.Fact, error) {
	var fact = engine.NilFact
	for _, path := range collector.Search(".*_test.go") {
		re := regexp.MustCompile(`func\sTest.*(Integration)\(t\s\*testing.T\)\s`)
		contents, err := os.ReadFile(path)
		if err != nil {
			return fact, err
		}
		if re.Match(contents) {
			fact = GolangIntegrationTestsExistsFact
		}
	}
	return fact, nil
}

var GolangCmdExistsRule engine.Rule = func(source string, _ engine.Collector, _ mapset.Set[engine.Fact]) (engine.Fact, error) {
	var fact = engine.NilFact
	stat, err := os.Stat(filepath.Join(source, "./cmd"))
	if os.IsNotExist(err) {
		return fact, nil
	} else if err != nil {
		return fact, err
	}
	if !stat.IsDir() {
		return fact, nil
	}
	fact = GolangCmdExistsFact
	return fact, nil
}

func init() {
	engine.AddToRuleset(&GoModExistsRule, &GolangTestExistsRule)
	engine.AddToRuleset(&GoModExistsRule, &GolangCmdExistsRule)
	engine.AddToRuleset(&GolangTestExistsRule, &GolangIntegrationTestExistsRule)
}
