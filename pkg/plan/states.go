package plan

type State int

const (
	OnDemandState State = iota
	FeebackState
	StageState
	QAState
	ReleaseState
)
