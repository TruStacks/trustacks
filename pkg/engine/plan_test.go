package engine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestActionPlanAddAction(t *testing.T) {
	ap := NewActionPlan()
	ap.AddAction("test")
	assert.Equal(t, "test", ap.Actions[0])
}

func TestCheckInputs(t *testing.T) {
	var mockInput InputField = "TEST"
	var previousRegisteredActions = registeredActions
	defer func() {
		registeredActions = previousRegisteredActions
	}()
	mockActionA := &Action{Name: "actionA", Inputs: []InputField{mockInput}}
	registeredActions = map[string]*Action{"actionA": mockActionA}
	ap := NewActionPlan()
	ap.Actions = []string{"actionA"}
	t.Run("withInputDefined", func(t *testing.T) {
		t.Setenv("TEST", "test")
		assert.NoError(t, ap.checkInputs([]string{"actionA"}))
	})
	t.Run("withoutInputDefined", func(t *testing.T) {
		assert.Error(t, ap.checkInputs([]string{"actionA"}))
	})
}

func TestStageActions(t *testing.T) {
	var previousRegisteredActions = registeredActions
	defer func() {
		registeredActions = previousRegisteredActions
	}()
	mockActionA := &Action{Name: "actionA", Stage: CommitStage}
	mockActionB := &Action{Name: "actionB", Stage: CommitStage}
	mockActionC := &Action{Name: "actionC", Stage: DeployStage}
	registeredActions = map[string]*Action{
		"actionA": mockActionA,
		"actionB": mockActionB,
		"actionC": mockActionC,
	}
	ap := NewActionPlan()
	ap.Actions = []string{"actionA", "actionB", "actionC"}
	actions := ap.stageActions([]string{actionStages[CommitStage]})
	assert.Contains(t, actions, "actionA")
	assert.Contains(t, actions, "actionB")
	assert.NotContains(t, actions, "actionC")
}
