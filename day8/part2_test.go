package day8

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
	a.Analyze()
	t.Log("\n" + a.String())
}

func TestPart2Example(t *testing.T) {
	entries, err := ParseEntries(util.Lines(exampleInput))
	assert.NoError(t, err)
	for i, e := range entries {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			a := NewAnalysis(e)
			a.Analyze()
			t.Log("\n" + a.String())
		})
	}
}
