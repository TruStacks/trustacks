package engine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewFact(t *testing.T) {
	assert.Equal(t, Fact(1), NewFact())
	assert.Equal(t, Fact(2), NewFact())
	assert.Equal(t, Fact(3), NewFact())
}
