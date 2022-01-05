package day14

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1Example(t *testing.T) {
	polymer, rules, err := Parse(exampleInput)
	assert.NoError(t, err)
	assert.Equal(t, "NNCB", polymer)
	for i := 0; i < 10; i++ {
		polymer = Apply(polymer, rules)
		switch i {
		case 0:
			assert.Equal(t, "NCNBCHB", polymer)
		case 1:
			assert.Equal(t, "NBCCNBBBCBHCB", polymer)
		case 2:
			assert.Equal(t, "NBBBCNCCNBBNBNBBCHBHHBCHB", polymer)
		case 3:
			assert.Equal(t, "NBBNBNBBCCNBCNCCNBBNBBNBBBNBBNBBCBHCBHHNHCBBCBHCB", polymer)
		}
	}
	stats := Analyze(polymer)
	assert.Equal(t, []Stat{{'H', 161}, {'C', 298}, {'N', 865}, {'B', 1749}}, stats)
}

func TestPart1Challenge(t *testing.T) {
	polymer, rules, err := Parse(challengeInput)
	assert.NoError(t, err)
	for i := 0; i < 10; i++ {
		polymer = Apply(polymer, rules)
	}
	stats := Analyze(polymer)
	score := stats[len(stats)-1].Count - stats[0].Count
	t.Logf("challenge score = %d", score)
}
