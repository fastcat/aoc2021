package day14

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func assertAnalyze2(t *testing.T, polymer string, counts map[rune]int64) {
	t.Helper()
	stats := Analyze2(polymer, Pairs(polymer))
	assert.Len(t, stats, len(counts))
	for _, s := range stats {
		assert.Equal(t, counts[s.Element], s.Count, string(s.Element))
	}
}

func TestAnalyze2(t *testing.T) {
	assertAnalyze2(t, "NNCB", map[rune]int64{'N': 2, 'C': 1, 'B': 1})
}

func TestPart2Example(t *testing.T) {
	polymer, rules, err := Parse(exampleInput)
	assert.NoError(t, err)
	counts := Pairs(polymer)
	for i := 0; i < 40; i++ {
		Apply2(counts, rules)
	}
	stats := Analyze2(polymer, counts)
	assert.Equal(t, Stat{'H', 3849876073}, stats[0])
	assert.Equal(t, Stat{'B', 2192039569602}, stats[len(stats)-1])
}

func TestPart2Challenge(t *testing.T) {
	polymer, rules, err := Parse(challengeInput)
	assert.NoError(t, err)
	counts := Pairs(polymer)
	for i := 0; i < 40; i++ {
		Apply2(counts, rules)
	}
	stats := Analyze2(polymer, counts)
	score := stats[len(stats)-1].Count - stats[0].Count
	t.Logf("challenge score = %d", score)
}
