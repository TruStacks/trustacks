package engine

type Stage int

const (
	OnDemand Stage = iota
	CommitStage
	AcceptanceStage
	NonFunctionalStage
	DeployStage
	ReleaseStage
)

var actionStages = []string{
	"", // DO NOT DELETE!
	"commit",
	"acceptance",
	"nonfunctional",
	"deploy",
	"release",
}

func GetStage(stage Stage) string {
	return actionStages[stage]
}
