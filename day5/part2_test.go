package day5

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestExample2(t *testing.T) {
	lines, err := ParseLines(exampleInput)
	require.NoError(t, err)
	grid := GridOverlaps(lines, true)
	const threshold = 2
	nAbove := NumPointsAtLeast(grid, threshold, t.Logf)
	assert.Equal(t, 12, nAbove)
}

func TestChallenge2(t *testing.T) {
	lines, err := ParseLines(challengeInput)
	require.NoError(t, err)
	grid := GridOverlaps(lines, true)
	const threshold = 2
	nAbove := NumPointsAtLeast(grid, threshold, nil)
	t.Logf("found %d points >= %d", nAbove, threshold)
}
