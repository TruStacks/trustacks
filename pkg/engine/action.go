package engine

import "dagger.io/dagger"

var registeredActions = map[string]*Action{}

type Action struct {
	Name                   string
	DisplayName            string
	Description            string
	Image                  func(*Config) string
	Stage                  Stage
	Script                 func(*dagger.Container, map[string]interface{}, *ActionUtilities) error
	InputArtifacts         []Artifact
	OptionalInputArtifacts []Artifact
	OutputArtifacts        []Artifact
	Caches                 []string
	Inputs                 []InputField
	AdmissionCriteria      []Fact
	ExclusionCriteria      []Fact
}

func GetAction(name string) *Action {
	for _, action := range registeredActions {
		if action.Name == name {
			return action
		}
	}
	return nil
}

func RegisterAction(action *Action) {
	registeredActions[action.Name] = action
}
