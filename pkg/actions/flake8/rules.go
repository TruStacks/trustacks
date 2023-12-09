package flake8

import (
	"os"
	"path/filepath"

	"github.com/bigkevmcd/go-configparser"
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/trustacks/trustacks/pkg/actions/python"
	"github.com/trustacks/trustacks/pkg/engine"
)

var (
	Flake8ConfigExistsFact = engine.NewFact()
)

var Flake8ConfigExistsRule engine.Rule = func(source string, collector engine.Collector, _ mapset.Set[engine.Fact]) (engine.Fact, error) {
	var fact = engine.NilFact
	var flake8ConfigFiles = []string{
		".flake8",
		"setup.cfg",
		"tox.ini",
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
		for _, config := range flake8ConfigFiles {
			if config == entry.Name() {
				cfg, err := configparser.NewConfigParserFromFile(filepath.Join(source, entry.Name()))
				if err != nil {
					return fact, err
				}
				if cfg.HasSection("flake8") {
					fact = Flake8ConfigExistsFact
					matched = true
				}
			}
		}
	}
	return fact, nil
}

func init() {
	engine.AddToRuleset(&python.PyProjectTomlExistsRule, &Flake8ConfigExistsRule)
}
