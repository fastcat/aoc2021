package day10

import (
	"testing"

	"github.com/fastcat/aoc2021/util"
	"github.com/stretchr/testify/assert"
)

func TestPart2LineScore(t *testing.T) {
	for _, tt := range []struct {
		input string
		want  int
	}{
		{"[({(<(())[]>[[{[]{<()<>>", 288957},
		{"[(()[<>])]({[<{<<[]>>(", 5566},
		{"(((({<>}<{<{<>}{[]{[]{}", 1480781},
		{"{<[[]]>}<{[{[{[]{()[[[]", 995444},
		{"<{([{{}}[<[[[<>{}]]]>[]]", 294},
	} {
		t.Run(tt.input, func(t *testing.T) {
			ok, got := Part2LineScore(tt.input)
			assert.True(t, ok)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestPart2Example(t *testing.T) {
	lines := util.Lines(exampleInput)
	score := Part2TotalScore(lines...)
	assert.Equal(t, 288957, score)
}

func TestPart2Challenge(t *testing.T) {
	lines := util.Lines(challengeInput)
	score := Part2TotalScore(lines...)
	t.Logf("challenge score = %d", score)
}
