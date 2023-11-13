package engine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetInput(t *testing.T) {
	input := GetInput(string(ContainerRegistry))
	assert.Equal(t, inputs["CONTAINER_REGISTRY"], input)
}
