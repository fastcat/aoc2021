package day10

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseChunkHappy(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  ChunkList
	}{
		{"simpleP", "()", Chunks(ChunkP())},
		{"simpleS", "[]", Chunks(ChunkS())},
		{"simpleC", "{}", Chunks(ChunkC())},
		{"simpleA", "<>", Chunks(ChunkA())},
		{"nestedP_PSCA", "([]{}<>)", Chunks(ChunkP(ChunkS(), ChunkC(), ChunkA()))},
		{"deep", "[<>({}){}[([])<>]]", Chunks(ChunkS(
			ChunkA(),
			ChunkP(ChunkC()),
			ChunkC(),
			ChunkS(ChunkP(ChunkS()), ChunkA()),
		))},
		{"adjacent", "()()", Chunks(ChunkP(), ChunkP())},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got ChunkList
			assert.NoError(t, chunksParser.ParseString(tt.name, tt.input, &got))
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestParseChunkIncomplete(t *testing.T) {
	for _, tt := range []struct {
		input string
		pos   int
	}{
		{"((())", 6},
		{"(", 2},
		{"((", 3},
		{"(()(", 5},
		{"({[<{<<[]>>(", 13},
		{"[(()[<>])]({[<{<<[]>>(", 23},
	} {
		t.Run(tt.input, func(t *testing.T) {
			var got ChunkList
			err := chunksParser.ParseString("incomplete.chunk", tt.input, &got)
			// t.Logf("%v", err)
			assert.Error(t, err)
			isInc, pe := IsIncomplete(err)
			assert.True(t, isInc)
			isInv, _ := IsInvalid(err)
			assert.False(t, isInv)
			assert.Equal(t, tt.pos, pe.Position().Column)
			assert.Equal(t, tt.input, got.String())
			// t.Logf("%#v", got)
		})
	}
}

func TestParseChunkInvalid(t *testing.T) {
	tests := []struct {
		input string
		pos   int
	}{
		{"(]", 2},
		{"{()()()>", 8},
		{"(((()))}", 8},
		{"<([]){()}[{}])", 14},
		{"((])", 3},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			var got ChunkList
			err := chunksParser.ParseString("invalid.chunk", tt.input, &got)
			// t.Logf("%v", err)
			assert.Error(t, err)
			isInc, _ := IsIncomplete(err)
			assert.False(t, isInc)
			isInv, pe := IsInvalid(err)
			assert.True(t, isInv)
			assert.Equal(t, tt.pos, pe.Position().Column)
			// t.Logf("%#v", got)
		})
	}
}
