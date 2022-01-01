package day11

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		want      Board
		assertion assert.ErrorAssertionFunc
	}{
		{
			"1x1",
			"1",
			Board{[]int{1}, 1, 1},
			nil,
		},
		{
			"2x2",
			"12\n34\n",
			Board{[]int{1, 2, 3, 4}, 2, 2},
			nil,
		},
		{
			"bogus",
			"12\n3a",
			Board{[]int{1, 2, 3, 0}, 2, 2},
			assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.input)
			if tt.assertion == nil {
				tt.assertion = assert.NoError
			}
			tt.assertion(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestBoard_Step(t *testing.T) {
	tests := []struct {
		name        string
		before      Board
		wantFlashes int
		after       Board
	}{
		{
			"1x1",
			MustParse("1"),
			0,
			MustParse("2"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Board{}
			*b = tt.before
			assert.Equal(t, tt.wantFlashes, b.Step())
			assert.Equal(t, tt.after, *b)
		})
	}
}
