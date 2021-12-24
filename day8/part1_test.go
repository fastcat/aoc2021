package day8

import (
	"testing"

	"github.com/fastcat/aoc2021/util"
	"github.com/stretchr/testify/assert"
)

func countSimple(entries ...Entry) int {
	n := 0
	for _, e := range entries {
		a := NewAnalysis(e)
		a.Rule01BitLengths()
		for _, o := range e.Outputs {
			v := a.Decode(o)
			if v.Len() == 1 {
				n++
			}
		}
	}
	return n
}

func TestPart1Example(t *testing.T) {
	entries, err := ParseEntries(util.Lines(exampleInput))
	assert.NoError(t, err)
	n := countSimple(entries...)
	assert.Equal(t, 26, n)
}

func TestPart1Challenge(t *testing.T) {
	entries, err := ParseEntries(util.Lines(challengeInput))
	assert.NoError(t, err)
	n := countSimple(entries...)
	t.Logf("have %d simple instances of 1,4,7,8", n)
}
