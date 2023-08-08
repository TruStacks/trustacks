package plan

type Stage int

const (
	OnDemandStage Stage = iota
	FeedbackStage
	PackageStage
	StageStage
	QAStage
	ReleaseStage
)

var actionStages = []string{
	"", // DO NOT DELETE!
	"feedback",
	"package",
	"stage",
	"qa",
	"release",
}
