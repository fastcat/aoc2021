package day18

import (
	"math"
	"testing"

	"github.com/fastcat/aoc2021/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func maxPair(t *testing.T, input string) int {
	lines := util.Lines(input)
	n := make([]*Node, len(lines))
	var err error
	for i, l := range lines {
		n[i], err = Parse(l)
		require.NoError(t, err)
	}

	max := math.MinInt
	// is there a better way than brute force here?
	for i := 0; i < len(n); i++ {
		for j := i + 1; j < len(n); j++ {
			s := Add(n[i], n[j])
			if m := s.Magnitude(); m > max {
				max = m
			}
			s = Add(n[j], n[i])
			if m := s.Magnitude(); m > max {
				max = m
			}
		}
	}
	return max
}

func TestPart2Example(t *testing.T) {
	m := maxPair(t, exampleInput)
	assert.Equal(t, 3993, m)
}

func TestPart2Challenge(t *testing.T) {
	m := maxPair(t, challengeInput)
	t.Logf("challenge max pair = %d", m)
}
