package engine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetStage(t *testing.T) {
	assert.Equal(t, GetStage(FeedbackStage), "feedback")
}
