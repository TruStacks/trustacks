package engine

import (
	"testing"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/stretchr/testify/assert"
)

func TestSchedulerAssignActivityStage(t *testing.T) {
	defer func() {
		registeredActions = map[string]*Action{}
	}()
	var mockArtifact Artifact = 1
	actionA := &Action{Name: "actionA", Stage: OnDemandStage, OutputArtifacts: []Artifact{mockArtifact}}
	actionB := &Action{Name: "actionB", Stage: FeedbackStage}
	actionC := &Action{Name: "actionC", Stage: PreleaseStage, InputArtifacts: []Artifact{mockArtifact}}
	registeredActions = map[string]*Action{
		"actionA": actionA,
		"actionB": actionB,
		"actionC": actionC,
	}
	s := newScheduler()
	assignments := s.assignActivityStage([]string{"actionA", "actionB", "actionC"})
	assert.True(t, assignments[OnDemandStage].Contains(actionA))
	assert.True(t, assignments[FeedbackStage].Contains(actionB))
	assert.True(t, assignments[PreleaseStage].Contains(actionC))
}

func TestBindActionInputs(t *testing.T) {
	var mockArtifact Artifact = 1
	assignments := map[Stage]mapset.Set[*Action]{
		OnDemandStage: mapset.NewSet[*Action](),
		FeedbackStage: mapset.NewSet[*Action](),
		PreleaseStage: mapset.NewSet[*Action](),
	}
	assignments[OnDemandStage].Add(&Action{Name: "actionA", Stage: OnDemandStage, OutputArtifacts: []Artifact{mockArtifact}})
	assignments[FeedbackStage].Add(&Action{Name: "actionB", Stage: FeedbackStage})
	assignments[PreleaseStage].Add(&Action{Name: "actionC", Stage: PreleaseStage, InputArtifacts: []Artifact{mockArtifact}})
	s := newScheduler()
	s.bindActionInputs(assignments)
	assert.True(t, s.requiredInputs[mockArtifact].Contains(PreleaseStage))
}

func TestFindOptionalInputOccurances(t *testing.T) {
	var mockArtifact Artifact = 1
	assignments := map[Stage]mapset.Set[*Action]{
		OnDemandStage: mapset.NewSet[*Action](),
		FeedbackStage: mapset.NewSet[*Action](),
		PreleaseStage: mapset.NewSet[*Action](),
	}
	assignments[OnDemandStage].Add(&Action{Name: "actionA", Stage: OnDemandStage, OutputArtifacts: []Artifact{mockArtifact}})
	assignments[FeedbackStage].Add(&Action{Name: "actionB", Stage: FeedbackStage, OptionalInputArtifacts: []Artifact{mockArtifact}})
	assignments[PreleaseStage].Add(&Action{Name: "actionC", Stage: PreleaseStage, InputArtifacts: []Artifact{mockArtifact}})
	s := newScheduler()
	s.bindActionInputs(assignments)
	assert.True(t, s.optionalInputs[mockArtifact].Contains(FeedbackStage))
	assert.True(t, s.requiredInputs[mockArtifact].Contains(PreleaseStage))
}

func TestSchedulerAssigneOnDemandActions(t *testing.T) {
	var mockArtifact Artifact = 1
	actionA := &Action{Name: "actionA", Stage: OnDemandStage, OutputArtifacts: []Artifact{mockArtifact}}
	assignments := map[Stage]mapset.Set[*Action]{
		OnDemandStage: mapset.NewSet[*Action](),
		FeedbackStage: mapset.NewSet[*Action](),
	}
	assignments[OnDemandStage].Add(actionA)
	assert.True(t, assignments[OnDemandStage].Contains(actionA))
	s := newScheduler()
	s.requiredInputs[mockArtifact] = mapset.NewSet[Stage](ReleaseStage, FeedbackStage)
	if err := s.assignOnDemandActions(assignments); err != nil {
		t.Fatal(err)
	}
	assert.False(t, assignments[OnDemandStage].Contains(actionA))
	assert.True(t, assignments[FeedbackStage].Contains(actionA))
}

func TestSchedulerSortActions(t *testing.T) {
	t.Run("resolvableInputs", func(t *testing.T) {
		var mockArtifactA Artifact = 1
		var mockArtifactB Artifact = 2
		actionA := &Action{Name: "actionA", Stage: FeedbackStage, OutputArtifacts: []Artifact{mockArtifactA}}
		actionB := &Action{Name: "actionB", Stage: FeedbackStage, OutputArtifacts: []Artifact{mockArtifactB}, InputArtifacts: []Artifact{mockArtifactA}}
		actionC := &Action{Name: "actionC", Stage: FeedbackStage, InputArtifacts: []Artifact{mockArtifactB}}
		assignments := map[Stage]mapset.Set[*Action]{
			FeedbackStage: mapset.NewSet[*Action](actionB, actionC, actionA),
		}
		s := newScheduler()
		sortedAssignments, err := s.sortActions(assignments)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, sortedAssignments[FeedbackStage][0], actionA)
		assert.Equal(t, sortedAssignments[FeedbackStage][1], actionB)
		assert.Equal(t, sortedAssignments[FeedbackStage][2], actionC)
	})
	t.Run("unresolvableInputs", func(t *testing.T) {
		var mockArtifactA Artifact = 1
		var mockArtifactB Artifact = 2
		actionA := &Action{Name: "actionA", Stage: FeedbackStage, OutputArtifacts: []Artifact{mockArtifactA}}
		actionB := &Action{Name: "actionB", Stage: FeedbackStage, InputArtifacts: []Artifact{mockArtifactB}}
		assignments := map[Stage]mapset.Set[*Action]{
			FeedbackStage: mapset.NewSet[*Action](actionB, actionA),
		}
		s := newScheduler()
		_, err := s.sortActions(assignments)
		assert.ErrorContains(t, err, "the scheduler has detected unresolved inputs for the following actions: 'actionB'")
	})
}

func TestSchedulerSchedule(t *testing.T) {
	defer func() {
		registeredActions = map[string]*Action{}
	}()
	var mockArtifactA Artifact = 1
	var mockArtifactB Artifact = 2
	actionA := &Action{Name: "actionA", Stage: OnDemandStage, OutputArtifacts: []Artifact{mockArtifactA}}
	actionB := &Action{Name: "actionB", Stage: OnDemandStage, OutputArtifacts: []Artifact{mockArtifactB}}
	actionC := &Action{Name: "actionC", Stage: FeedbackStage}
	actionD := &Action{Name: "actionD", Stage: PreleaseStage, InputArtifacts: []Artifact{mockArtifactA}}
	actionE := &Action{Name: "actionE", Stage: ReleaseStage, InputArtifacts: []Artifact{mockArtifactB}}
	registeredActions = map[string]*Action{
		"actionA": actionA,
		"actionB": actionB,
		"actionC": actionC,
		"actionD": actionD,
		"actionE": actionE,
	}
	s := newScheduler()
	schedule, err := s.schedule([]string{"actionA", "actionB", "actionC", "actionD", "actionE"})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, schedule[FeedbackStage][0], actionC)
	assert.Equal(t, schedule[PreleaseStage][0], actionA)
	assert.Equal(t, schedule[PreleaseStage][1], actionD)
	assert.Equal(t, schedule[ReleaseStage][0], actionB)
	assert.Equal(t, schedule[ReleaseStage][1], actionE)
}
