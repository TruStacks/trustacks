package engine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAction(t *testing.T) {
	defer func() {
		registeredActions = map[string]*Action{}
	}()
	registeredActions = map[string]*Action{
		"test": {Name: "test"},
	}
	action := GetAction("test")
	assert.Equal(t, "test", action.Name)
}

func TestRegisterAction(t *testing.T) {
	defer func() {
		registeredActions = map[string]*Action{}
	}()
	assert.Len(t, registeredActions, 0)
	RegisterAction(&Action{Name: "test"})
	assert.Len(t, registeredActions, 1)
}
