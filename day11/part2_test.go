package day11

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart2Example(t *testing.T) {
	b, err := Parse(exampleInput)
	assert.NoError(t, err)
	for i := 1; i <= 195; i++ {
		b.Step()
	}
	assert.True(t, b.AllZero())
}

func TestPart2Challenge(t *testing.T) {
	b, err := Parse(challengeInput)
	assert.NoError(t, err)
	for i := 1; ; i++ {
		b.Step()
		if b.AllZero() {
			t.Logf("challenge synchronized to zeros after %d steps", i)
			break
		}
	}
}
