package day09

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart2Example(t *testing.T) {
	b, err := ParseBoard(exampleInput)
	assert.NoError(t, err)
	basins := b.Basins()
	assert.ElementsMatch(t, []Basin{
		{Point{0, 1}, 3},
		{Point{0, 9}, 9},
		{Point{2, 2}, 14},
		{Point{4, 6}, 9},
	}, basins)
	score := Part2BasinScore(basins)
	assert.Equal(t, 1134, score)
}

func TestPart2Challenge(t *testing.T) {
	b, err := ParseBoard(challengeInput)
	assert.NoError(t, err)
	basins := b.Basins()
	score := Part2BasinScore(basins)
	t.Logf("challeng score = %d", score)
}
