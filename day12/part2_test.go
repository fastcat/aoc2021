package day12

import (
	"strconv"
	"testing"

	"github.com/fastcat/aoc2021/util"
	"github.com/stretchr/testify/assert"
)

func TestPart2Examples(t *testing.T) {
	for i, input := range examplesInputStanzas {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			g := MustParse(input)
			want := util.Lines(examplePaths2[i])
			var got []string
			visitor := func(p Path) {
				got = append(got, p.String())
			}
			count := g.CountPaths2(visitor)
			if len(want) == 1 {
				wantCount, err := strconv.Atoi(want[0])
				assert.NoError(t, err)
				assert.Equal(t, wantCount, count)
			} else {
				assert.Equal(t, len(want), count)
				assert.ElementsMatch(t, want, got)
			}
		})
	}
}

func TestPart2Challenge(t *testing.T) {
	g := MustParse(challengeInput)
	count := g.CountPaths2(nil)
	t.Logf("challenge path count = %d", count)
}
