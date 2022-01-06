package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseIntBoardCompact(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		want      *IntBoard
		assertion assert.ErrorAssertionFunc
	}{
		{
			"1x1",
			"1",
			&IntBoard{[]int{1}, 1, 1},
			nil,
		},
		{
			"small",
			"123\n456\n789\n012",
			&IntBoard{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2}, 3, 4},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseIntBoardCompact(tt.input)
			if tt.assertion == nil {
				tt.assertion = assert.NoError
			}
			tt.assertion(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestIntBoardIterator_NextRC(t *testing.T) {
	board := &IntBoard{
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		3, 3,
	}
	tests := []struct {
		name      string
		before    IntBoardIterator
		want      IntBoardIterator
		wantValid bool
		wantValue int
	}{
		{
			"0,0",
			board.IteratorAt(0, 0),
			board.IteratorAt(0, 1),
			true,
			2,
		},
		{
			"2,0",
			board.IteratorAt(0, 2),
			board.IteratorAt(1, 0),
			true,
			4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.True(t, tt.before.Valid())
			after := tt.before.NextRC()
			assert.Equal(t, tt.want, after)
			if assert.Equal(t, tt.wantValid, after.Valid()) && tt.wantValid {
				assert.Equal(t, tt.wantValue, after.Value())
			}
		})
	}
}
