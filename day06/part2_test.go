package day06

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample2(t *testing.T) {
	initial := InitState(exampleInput)
	day := 0
	state := initial
	for ; day < 256; day++ {
		state = Grow(state)
	}
	assert.EqualValues(t, 26984457539, Total(state))
}

func TestChallenge2(t *testing.T) {
	initial := InitState(challengeInput)
	day := 0
	state := initial
	for ; day < 256; day++ {
		state = Grow(state)
	}
	t.Logf("After %d days: %v (%d)", day, state, Total(state))
}
