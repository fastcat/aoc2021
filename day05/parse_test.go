package day05

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParsePoints(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		want      []Line
		assertion assert.ErrorAssertionFunc
	}{
		{
			"example 1",
			"0,9 -> 5,9",
			[]Line{{Point{0, 9}, Point{5, 9}}},
			assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseLines(tt.input)
			tt.assertion(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
