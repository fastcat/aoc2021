package day6

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	initial := InitState(exampleInput)
	assert.Equal(t, GrowthState{0, 1, 1, 2, 1}, initial)
	assert.Equal(t, len(exampleInput), Total(initial))
	day := 0
	state := initial
	for ; day < 18; day++ {
		state = Grow(state)
		t.Logf("After %d days: %v (%d)", day+1, state, Total(state))
	}
	assert.Equal(t, 26, Total(state))
	for ; day < 80; day++ {
		state = Grow(state)
		// t.Logf("After %d days: %v (%d)", day+1, state, Total(state))
	}
	assert.Equal(t, 5934, Total(state))
}

func TestChallenge1(t *testing.T) {
	initial := InitState(challengeInput)
	day := 0
	state := initial
	for ; day < 80; day++ {
		state = Grow(state)
		// t.Logf("After %d days: %v (%d)", day+1, state, Total(state))
	}
	t.Logf("After %d days: %v (%d)", day+1, state, Total(state))
}
