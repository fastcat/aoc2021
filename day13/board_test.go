package day13

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseBoard(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  []Point
	}{
		{
			"trivial",
			"0,0",
			[]Point{{0, 0}},
		},
		{
			"sparse",
			"0,1\n99,100",
			[]Point{{0, 1}, {99, 100}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseBoard(tt.input)
			assert.NoError(t, err)
			gotCells := make([]Point, 0, len(tt.want))
			for pt, val := range got.cells {
				if assert.True(t, val, pt) {
					gotCells = append(gotCells, pt)
				}
			}
			assert.ElementsMatch(t, tt.want, gotCells)
		})
	}
}

func TestBoard_String(t *testing.T) {
	tests := []struct {
		name  string
		cells []Point
		want  string
	}{
		{"trivial", []Point{{0, 0}}, "#"},
		{
			"diagonal",
			[]Point{{0, 0}, {3, 3}},
			"#...\n....\n....\n...#",
		},
		{
			"zigzag",
			[]Point{{0, 0}, {2, 1}, {1, 2}, {3, 3}},
			"#...\n..#.\n.#..\n...#",
		},
		{
			"multi",
			[]Point{{0, 0}, {1, 0}, {2, 0}},
			"###",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cells := make(map[Point]bool, len(tt.cells))
			for _, pt := range tt.cells {
				cells[pt] = true
			}
			b := &Board{cells: cells}
			assert.Equal(t, tt.want, b.String())
		})
	}
}

func TestParseFolds(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  []Fold
	}{
		{
			"x=3",
			"fold along x=3",
			[]Fold{{Pos: 3, Horizontal: false}},
		},
		{
			"x=6,y=2",
			"fold along x=6\nfold along y=2",
			[]Fold{
				{Pos: 6, Horizontal: false},
				{Pos: 2, Horizontal: true},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseFolds(tt.input)
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestBoard_ApplyFold(t *testing.T) {
	tests := []struct {
		name   string
		before []Point
		fold   Fold
		after  []Point
	}{
		{
			"simple horiz",
			[]Point{{2, 2}},
			Fold{Pos: 1, Horizontal: true},
			[]Point{{2, 0}},
		},
		{
			"simple vert",
			[]Point{{2, 2}},
			Fold{Pos: 1, Horizontal: false},
			[]Point{{0, 2}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Board{cells: make(map[Point]bool, len(tt.before))}
			for _, pt := range tt.before {
				b.cells[pt] = true
			}
			b.ApplyFold(tt.fold)
			after := make([]Point, 0, len(b.cells))
			for pt, val := range b.cells {
				if assert.True(t, val, pt) {
					after = append(after, pt)
				}
			}
			assert.ElementsMatch(t, tt.after, after)
		})
	}
}
