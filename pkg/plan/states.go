package plan

type State int

const (
	OnDemandState State = iota
	FeebackState
	PackageState
	StageState
	QAState
	ReleaseState
)
