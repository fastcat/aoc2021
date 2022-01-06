package day15

import (
	"testing"

	"github.com/fastcat/aoc2021/util"
	"github.com/stretchr/testify/assert"
)

func TestPart1Example(t *testing.T) {
	b, err := util.ParseIntBoardCompact(exampleInput)
	assert.NoError(t, err)
	p := PathTo(b, util.Point{R: b.Height - 1, C: b.Width - 1})
	p.Fill()
	cost := p.CostFrom(util.Point{R: 0, C: 0})
	assert.Equal(t, 40, cost)
}

func TestPart1Challenge(t *testing.T) {
	b, err := util.ParseIntBoardCompact(challengeInput)
	assert.NoError(t, err)
	p := PathTo(b, util.Point{R: b.Height - 1, C: b.Width - 1})
	p.Fill()
	cost := p.CostFrom(util.Point{R: 0, C: 0})
	t.Logf("challenge best cost = %d", cost)
}
