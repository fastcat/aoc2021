package day10

import (
	"fmt"
	"testing"

	"github.com/fastcat/aoc2021/util"
	"github.com/stretchr/testify/assert"
)

func TestPart1Example(t *testing.T) {
	lines := util.Lines(exampleInput)
	score := 0
	for i, l := range lines {
		var c ChunkList
		err := chunksParser.ParseString(fmt.Sprintf("example/%d", i), l, &c)
		score += Part1Score(l, err)
	}
	assert.Equal(t, 26397, score)
}

func TestPart1Challenge(t *testing.T) {
	lines := util.Lines(challengeInput)
	score := 0
	for i, l := range lines {
		var c ChunkList
		err := chunksParser.ParseString(fmt.Sprintf("challenge/%d", i), l, &c)
		score += Part1Score(l, err)
	}
	t.Logf("challenge score = %d", score)
}
