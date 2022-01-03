package day12

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMustParse(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  Graph
	}{
		{
			"trivial",
			"start-end",
			Graph{
				Start: {End: true},
				End:   {Start: true},
			},
		},
		{
			"small",
			"start-a\na-b\nb-end",
			Graph{
				Start: {"a": true},
				"a":   {Start: true, "b": true},
				"b":   {"a": true, End: true},
				End:   {"b": true},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, MustParse(tt.input))
		})
	}
}
