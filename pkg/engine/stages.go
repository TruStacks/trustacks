package engine

type Stage int

const (
	OnDemandStage Stage = iota
	FeedbackStage
	PackageStage
	PreleaseStage
	ReleaseStage
)

var actionStages = []string{
	"", // DO NOT DELETE!
	"feedback",
	"package",
	"prerelease",
	"release",
}

func GetStage(stage Stage) string {
	return actionStages[stage]
}
