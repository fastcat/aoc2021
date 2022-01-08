package day18

import (
	_ "embed"
	"testing"

	"github.com/fastcat/aoc2021/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func sumLines(t *testing.T, input string) *Node {
	lines := util.Lines(input)
	var s *Node
	for _, l := range lines {
		t.Logf("adding %s", l)
		a, err := Parse(l)
		require.NoError(t, err)
		a.valid(t, nil)
		if s == nil {
			s = a
			continue
		}
		s = Add(s, a)
		t.Logf("got %s", s.String())
		s.valid(t, nil)
	}
	return s
}

//go:embed example.txt
var exampleInput string

func TestPart1Example(t *testing.T) {
	s := sumLines(t, exampleInput)
	assert.Equal(t, "[[[[6,6],[7,6]],[[7,7],[7,0]]],[[[7,7],[7,7]],[[7,8],[9,9]]]]", s.String())
	assert.Equal(t, 4140, s.Magnitude())
}

//go:embed challenge.txt
var challengeInput string

func TestPart1Challenge(t *testing.T) {
	s := sumLines(t, challengeInput)
	t.Logf("challenge sum = %s", s)
	t.Logf("challenge mag = %d", s.Magnitude())
}
