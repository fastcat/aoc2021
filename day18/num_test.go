package day18

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParse(t *testing.T) {
	tests := []struct {
		input string
		want  *Node
	}{
		{"[1,2]", N(1, 2)},
		{"[[1,2],3]", N(N(1, 2), 3)},
		{"[9,[8,7]]", N(9, N(8, 7))},
		// more from the examples
		// [[1,9],[8,5]]
		// [[[[1,2],[3,4]],[[5,6],[7,8]]],9]
		// [[[9,[3,8]],[[0,9],6]],[[[3,7],[4,9]],3]]
		// [[[[1,3],[5,3]],[[1,3],[8,7]]],[[[4,9],[6,9]],[[8,2],[7,3]]]]
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got, err := Parse(tt.input)
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestAdd(t *testing.T) {
	tests := []struct {
		l, r string
		want string
	}{
		{
			"[1,2]", "[3,4]",
			"[[1,2],[3,4]]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.l+"+"+tt.r, func(t *testing.T) {
			l, err := Parse(tt.l)
			require.NoError(t, err)
			r, err := Parse(tt.r)
			require.NoError(t, err)
			sum := Add(l, r)
			assert.Equal(t, tt.want, sum.String())
		})
	}
}

func TestNode_explodeWalk(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"[[[[[9,8],1],2],3],4]", "[[[[0,9],2],3],4]"},
		{"[7,[6,[5,[4,[3,2]]]]]", "[7,[6,[5,[7,0]]]]"},
		{"[[6,[5,[4,[3,2]]]],1]", "[[6,[5,[7,0]]],3]"},
		{"[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]", "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]"},
		{"[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]", "[[3,[2,[8,0]]],[9,[5,[7,0]]]]"},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			n, err := Parse(tt.input)
			require.NoError(t, err)
			exploded := n.explodeWalk(0)
			assert.Equal(t, tt.input != tt.want, exploded)
			assert.Equal(t, tt.want, n.String())
		})
	}
}
