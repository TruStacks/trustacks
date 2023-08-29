package rules

import (
	"os"
	"path/filepath"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/trustacks/pkg/engine"
)

var (
	SonarProjectPropertiesExists = engine.NewFact()
)

var sonarProjectPropertiesExistsRule engine.Rule = func(source string, collector engine.Collector, _ mapset.Set[engine.Fact]) (engine.Fact, error) {
	var fact = engine.NilFact
	if _, err := os.Stat(filepath.Join(source, "sonar-project.properties")); os.IsNotExist(err) {
		return fact, nil
	} else if err != nil {
		return fact, err
	}
	fact = SonarProjectPropertiesExists
	return fact, nil
}

func init() {
	engine.AddToRuleset(&expressDependencyExistsRule, &sonarProjectPropertiesExistsRule)
	engine.AddToRuleset(&reactScriptsBuildExsitsRule, &sonarProjectPropertiesExistsRule)

}
