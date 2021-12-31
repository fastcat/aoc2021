package day08

import (
	"fmt"
	"testing"

	"github.com/fastcat/aoc2021/util"
	"github.com/stretchr/testify/assert"
)

func TestPart2Demo(t *testing.T) {
	e, err := ParseEntry("acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab | cdfeb fcadb cdfeb cdbaf")
	assert.NoError(t, err)
	a := NewAnalysis(e)
	done := a.Analyze()
	t.Log("\n" + a.String())
	assert.True(t, done)
	assert.Equal(t, a.WireOptions, [7]Value{C, F, G, A, B, D, E})
	assert.Equal(t, a.PatternOptions, [10]DigitOption{Eight, Five, Two, Three, Seven, Nine, Six, Four, Zero, One})
}

func TestPart2ExampleAnalysis(t *testing.T) {
	entries, err := ParseEntries(util.Lines(exampleInput))
	assert.NoError(t, err)
	for i, e := range entries {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			a := NewAnalysis(e)
			done := a.Analyze()
			t.Log("\n" + a.String())
			assert.True(t, done)
		})
	}
}

func sumPart2(t *testing.T, input string) int {
	entries, err := ParseEntries(util.Lines(input))
	assert.NoError(t, err)
	sum := 0
	for _, e := range entries {
		a := NewAnalysis(e)
		assert.True(t, a.Analyze())
		digits := a.Decode(e.Outputs[:]...)
		value := 0
		exp := 1
		// digits are in reverse order here
		for i := len(digits) - 1; i >= 0; i-- {
			value += digits[i].Decode() * exp
			exp *= 10
		}
		sum += value
	}
	return sum
}
func TestPart2Example(t *testing.T) {
	sum := sumPart2(t, exampleInput)
	assert.Equal(t, 61229, sum)
}

func TestPart2Challenge(t *testing.T) {
	sum := sumPart2(t, challengeInput)
	t.Log(sum)
}
