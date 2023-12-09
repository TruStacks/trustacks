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
	actionA := &Action{Name: "actionA", Stage: OnDemand, OutputArtifacts: []Artifact{mockArtifact}}
	actionB := &Action{Name: "actionB", Stage: CommitStage}
	actionC := &Action{Name: "actionC", Stage: DeployStage, InputArtifacts: []Artifact{mockArtifact}}
	registeredActions = map[string]*Action{
		"actionA": actionA,
		"actionB": actionB,
		"actionC": actionC,
	}
	s := newScheduler()
	assignments := s.assignActivityStage([]string{"actionA", "actionB", "actionC"})
	assert.True(t, assignments[OnDemand].Contains(actionA))
	assert.True(t, assignments[CommitStage].Contains(actionB))
	assert.True(t, assignments[DeployStage].Contains(actionC))
}

func TestBindActionInputs(t *testing.T) {
	var mockArtifact Artifact = 1
	assignments := map[Stage]mapset.Set[*Action]{
		OnDemand:    mapset.NewSet[*Action](),
		CommitStage: mapset.NewSet[*Action](),
		DeployStage: mapset.NewSet[*Action](),
	}
	assignments[OnDemand].Add(&Action{Name: "actionA", Stage: OnDemand, OutputArtifacts: []Artifact{mockArtifact}})
	assignments[CommitStage].Add(&Action{Name: "actionB", Stage: CommitStage})
	assignments[DeployStage].Add(&Action{Name: "actionC", Stage: DeployStage, InputArtifacts: []Artifact{mockArtifact}})
	s := newScheduler()
	s.bindActionInputs(assignments)
	assert.True(t, s.requiredInputs[mockArtifact].Contains(DeployStage))
}

func TestFindOptionalInputOccurances(t *testing.T) {
	var mockArtifact Artifact = 1
	assignments := map[Stage]mapset.Set[*Action]{
		OnDemand:    mapset.NewSet[*Action](),
		CommitStage: mapset.NewSet[*Action](),
		DeployStage: mapset.NewSet[*Action](),
	}
	assignments[OnDemand].Add(&Action{Name: "actionA", Stage: OnDemand, OutputArtifacts: []Artifact{mockArtifact}})
	assignments[CommitStage].Add(&Action{Name: "actionB", Stage: CommitStage, OptionalInputArtifacts: []Artifact{mockArtifact}})
	assignments[DeployStage].Add(&Action{Name: "actionC", Stage: DeployStage, InputArtifacts: []Artifact{mockArtifact}})
	s := newScheduler()
	s.bindActionInputs(assignments)
	assert.True(t, s.optionalInputs[mockArtifact].Contains(CommitStage))
	assert.True(t, s.requiredInputs[mockArtifact].Contains(DeployStage))
}

func TestSchedulerAssignOnDemandActions(t *testing.T) {
	var mockArtifact Artifact = 1
	actionA := &Action{Name: "actionA", Stage: OnDemand, OutputArtifacts: []Artifact{mockArtifact}}
	assignments := map[Stage]mapset.Set[*Action]{
		OnDemand:    mapset.NewSet[*Action](),
		CommitStage: mapset.NewSet[*Action](),
	}
	assignments[OnDemand].Add(actionA)
	assert.True(t, assignments[OnDemand].Contains(actionA))
	s := newScheduler()
	t.Run("requiredInputs", func(t *testing.T) {
		s.requiredInputs[mockArtifact] = mapset.NewSet[Stage](ReleaseStage, CommitStage)
		if err := s.assignOnDemandActions(assignments); err != nil {
			t.Fatal(err)
		}
		assert.False(t, assignments[OnDemand].Contains(actionA))
		assert.True(t, assignments[CommitStage].Contains(actionA))
	})
	t.Run("optionalInputs", func(t *testing.T) {
		s.optionalInputs[mockArtifact] = mapset.NewSet[Stage](ReleaseStage, CommitStage)
		if err := s.assignOnDemandActions(assignments); err != nil {
			t.Fatal(err)
		}
		assert.False(t, assignments[OnDemand].Contains(actionA))
		assert.True(t, assignments[CommitStage].Contains(actionA))
	})
}

func TestSchedulerSortActions(t *testing.T) {
	t.Run("resolvableInputs", func(t *testing.T) {
		var mockArtifactA Artifact = 1
		var mockArtifactB Artifact = 2
		actionA := &Action{Name: "actionA", Stage: CommitStage, OutputArtifacts: []Artifact{mockArtifactA}}
		actionB := &Action{Name: "actionB", Stage: CommitStage, OutputArtifacts: []Artifact{mockArtifactB}, InputArtifacts: []Artifact{mockArtifactA}}
		actionC := &Action{Name: "actionC", Stage: CommitStage, InputArtifacts: []Artifact{mockArtifactB}}
		assignments := map[Stage]mapset.Set[*Action]{
			CommitStage: mapset.NewSet[*Action](actionB, actionC, actionA),
		}
		s := newScheduler()
		sortedAssignments, err := s.sortActions(assignments)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, sortedAssignments[CommitStage][0], actionA)
		assert.Equal(t, sortedAssignments[CommitStage][1], actionB)
		assert.Equal(t, sortedAssignments[CommitStage][2], actionC)
	})
	t.Run("unresolvableInputs", func(t *testing.T) {
		var mockArtifactA Artifact = 1
		var mockArtifactB Artifact = 2
		actionA := &Action{Name: "actionA", Stage: CommitStage, OutputArtifacts: []Artifact{mockArtifactA}}
		actionB := &Action{Name: "actionB", Stage: CommitStage, InputArtifacts: []Artifact{mockArtifactB}}
		assignments := map[Stage]mapset.Set[*Action]{
			CommitStage: mapset.NewSet[*Action](actionB, actionA),
		}
		s := newScheduler()
		_, err := s.sortActions(assignments)
		assert.ErrorContains(t, err, "the following action has inputs that cannot be resolved: 'actionB'")
	})
}

func TestSchedulerSchedule(t *testing.T) {
	defer func() {
		registeredActions = map[string]*Action{}
	}()
	var mockArtifactA Artifact = 1
	var mockArtifactB Artifact = 2
	actionA := &Action{Name: "actionA", Stage: OnDemand, OutputArtifacts: []Artifact{mockArtifactA}}
	actionB := &Action{Name: "actionB", Stage: OnDemand, OutputArtifacts: []Artifact{mockArtifactB}}
	actionC := &Action{Name: "actionC", Stage: CommitStage}
	actionD := &Action{Name: "actionD", Stage: DeployStage, InputArtifacts: []Artifact{mockArtifactA}}
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
	assert.Equal(t, schedule[CommitStage][0], actionC)
	assert.Equal(t, schedule[DeployStage][0], actionA)
	assert.Equal(t, schedule[DeployStage][1], actionD)
	assert.Equal(t, schedule[ReleaseStage][0], actionB)
	assert.Equal(t, schedule[ReleaseStage][1], actionE)
}
