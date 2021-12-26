package day9

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1Example(t *testing.T) {
	b, err := ParseBoard(exampleInput)
	assert.NoError(t, err)
	m := b.LocalMinima()
	assert.Equal(t, []Point{{0, 1}, {0, 9}, {2, 2}, {4, 6}}, m)
	s := Part1LocalMinimaRiskLevelSum(b)
	assert.Equal(t, 15, s)
}

func TestPart1Challenge(t *testing.T) {
	b, err := ParseBoard(challengeInput)
	assert.NoError(t, err)
	s := Part1LocalMinimaRiskLevelSum(b)
	t.Logf("challenge risk sum = %d", s)
}
