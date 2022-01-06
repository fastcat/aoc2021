package day15

import (
	"fmt"
	"strings"
	"testing"

	"github.com/fastcat/aoc2021/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func costMapStr(m *util.IntBoard) string {
	b := strings.Builder{}
	for i := m.Iterator(); i.Valid(); i = i.NextRC() {
		pt := i.Point()
		if pt.C == 0 && pt.R > 0 {
			b.WriteRune('\n')
		} else if pt.C > 0 {
			b.WriteRune(' ')
		}
		fmt.Fprintf(&b, "%3d", i.Value())
	}
	return b.String()
}

func TestPart2Example(t *testing.T) {
	b, err := util.ParseIntBoardCompact(exampleInput)
	assert.NoError(t, err)
	b = BuildFiveX(b)
	p := PathTo(b, util.Point{R: b.Height - 1, C: b.Width - 1})
	p.Fill()
	ps := PathTo(b, util.Point{R: b.Height - 1, C: b.Width - 1})
	slowMap := ps.fillSlow()
	fastMap := p.costMap()
	t.Logf("slowmap\n%s", costMapStr(slowMap))
	t.Logf("fastmap\n%s", costMapStr(fastMap))
	require.Equal(t, slowMap, fastMap)
	require.Equal(t, ps, p)
	cost := p.CostFrom(util.Point{R: 0, C: 0})
	assert.Equal(t, 315, cost)
}

func TestPart2Challenge(t *testing.T) {
	b, err := util.ParseIntBoardCompact(challengeInput)
	assert.NoError(t, err)
	b = BuildFiveX(b)
	p := PathTo(b, util.Point{R: b.Height - 1, C: b.Width - 1})
	p.Fill()
	cost := p.CostFrom(util.Point{R: 0, C: 0})
	t.Logf("challenge best cost = %d", cost)
}
