package day15

import (
	"testing"

	"github.com/fastcat/aoc2021/util"
	"github.com/stretchr/testify/assert"
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
