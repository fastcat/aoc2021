package day13

import (
	"fmt"
	"strings"

	"github.com/fastcat/aoc2021/util"
)

type Point struct {
	X, Y int
}

type Board struct {
	// board is pretty sparse, and we don't know the dims, so just use a set of
	// active cells
	cells map[Point]bool
}

func ParseBoard(input string) (*Board, error) {
	lines := util.Lines(input)
	board := &Board{cells: make(map[Point]bool, len(lines))}
	// first stanza is x,y pairs of set points
	for _, l := range lines {
		var pt Point
		if _, err := fmt.Sscanf(l, "%d,%d\n", &pt.X, &pt.Y); err != nil {
			return board, err
		}
		board.cells[pt] = true
	}
	return board, nil
}

func (b *Board) String() string {
	var lines [][]byte
	maxX := 0
	for pt, val := range b.cells {
		if !val {
			continue
		}
		if pt.X > maxX {
			maxX = pt.X
		}
		for len(lines) <= pt.Y {
			lines = append(lines, make([]byte, 0, maxX))
		}
		// TODO: this is inefficient
		for len(lines[pt.Y]) <= pt.X {
			lines[pt.Y] = append(lines[pt.Y], byte('.'))
		}
		lines[pt.Y][pt.X] = byte('#')
	}
	buf := strings.Builder{}
	for y, l := range lines {
		// TODO: this is inefficient
		for len(l) <= maxX {
			l = append(l, byte('.'))
		}
		if y > 0 {
			buf.WriteRune('\n')
		}
		buf.WriteString(string(l))
	}
	return buf.String()
}

type Fold struct {
	Pos        int
	Horizontal bool
}

func ParseFolds(input string) ([]Fold, error) {
	lines := util.Lines(input)
	folds := make([]Fold, 0, len(lines))
	for _, l := range lines {
		var f Fold
		var xy rune
		if _, err := fmt.Sscanf(l, "fold along %c=%d\n", &xy, &f.Pos); err != nil {
			return folds, err
		}
		switch xy {
		case 'x':
			// no-op: f.Horizontal = false
		case 'y':
			f.Horizontal = true
		default:
			return folds, fmt.Errorf("invalid fold axis '%c'", xy)
		}
		folds = append(folds, f)
	}
	return folds, nil
}

func ParseBoardAndFolds(input string) (*Board, []Fold, error) {
	stanzas := util.Stanzas(input)
	if len(stanzas) != 2 {
		return nil, nil, fmt.Errorf("expect 2 stanzas, got %d", len(stanzas))
	}
	b, err := ParseBoard(stanzas[0])
	if err != nil {
		return b, nil, err
	}
	f, err := ParseFolds(stanzas[1])
	return b, f, err
}

func (f Fold) String() string {
	xy := 'x'
	if f.Horizontal {
		xy = 'y'
	}
	return fmt.Sprintf("fold along %c=%d", xy, f.Pos)
}

func (b *Board) ApplyFold(fold Fold) {
	var predicate func(Point) bool
	var mapper func(Point) Point
	pp := 2 * fold.Pos
	if fold.Horizontal {
		predicate = func(p Point) bool { return p.Y > fold.Pos }
		mapper = func(p Point) Point { p.Y = pp - p.Y; return p }
	} else {
		predicate = func(p Point) bool { return p.X > fold.Pos }
		mapper = func(p Point) Point { p.X = pp - p.X; return p }
	}
	for pt, val := range b.cells {
		if !val {
			delete(b.cells, pt)
			continue
		}
		if predicate(pt) {
			pt2 := mapper(pt)
			b.cells[pt2] = true
			delete(b.cells, pt)
		}
	}
}

func (b *Board) ApplyFolds(folds []Fold) {
	for _, fold := range folds {
		b.ApplyFold(fold)
	}
}
