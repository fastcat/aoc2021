package day15

import (
	"testing"

	"github.com/fastcat/aoc2021/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseExampleInput(t *testing.T) {
	b, err := util.ParseIntBoardCompact(exampleInput)
	assert.NoError(t, err)
	assert.True(t, b.Valid())
	assert.Equal(t, 10, b.Width)
	assert.Equal(t, 10, b.Height)
	i := b.IteratorAt(3, 4)
	if assert.True(t, i.Valid()) {
		assert.Equal(t, 9, i.Value())
	}
}

func TestBuildFiveX(t *testing.T) {
	b, err := util.ParseIntBoardCompact("8")
	assert.NoError(t, err)
	b5 := BuildFiveX(b)
	assert.True(t, b5.Valid())
	require.Equal(t, &util.IntBoard{
		Values: []int{8, 9, 1, 2, 3, 9, 1, 2, 3, 4, 1, 2, 3, 4, 5, 2, 3, 4, 5, 6, 3, 4, 5, 6, 7},
		Width:  5,
		Height: 5,
	}, b5)
	b, err = util.ParseIntBoardCompact("12\n34")
	assert.NoError(t, err)
	b5 = BuildFiveX(b)
	assert.True(t, b5.Valid())
	want, err := util.ParseIntBoardCompact(
		"1223344556\n" +
			"3445566778\n" +
			"2334455667\n" +
			"4556677889\n" +
			"3445566778\n" +
			"5667788991\n" +
			"4556677889\n" +
			"6778899112\n" +
			"5667788991\n" +
			"7889911223\n",
	)
	assert.NoError(t, err)
	require.Equal(t, want, b5)
	b, err = util.ParseIntBoardCompact(exampleInput)
	assert.NoError(t, err)
	b5 = BuildFiveX(b)
	assert.True(t, b5.Valid())
	e5, err := util.ParseIntBoardCompact(exampleFiveXInput)
	assert.NoError(t, err)
	require.Equal(t, e5, b5)
}
