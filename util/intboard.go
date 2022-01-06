package util

import "fmt"

type IntBoard struct {
	Values        []int
	Width, Height int
}

func ParseIntBoardCompact(input string) (*IntBoard, error) {
	rows := Lines(input)
	board := &IntBoard{Height: len(rows)}
	for r, row := range rows {
		if board.Width != 0 {
			if board.Width != len(row) {
				return board, fmt.Errorf("row %d wrong width %d should be %d", r, len(row), board.Width)
			}
		} else {
			board.Width = len(row)
			board.Values = make([]int, board.Width*board.Height)
		}
		ro := board.Width * r
		for c := 0; c < len(row); c++ {
			value := int(row[c] - '0')
			if value < 0 || value > 9 {
				return board, fmt.Errorf("r,c %d,%d bad value %c=%d", r, c, row[c], value)
			}
			board.Values[ro+c] = value
		}
	}
	return board, nil
}

func (b *IntBoard) Valid() bool {
	return b != nil && b.Values != nil &&
		len(b.Values) == b.Width*b.Height
}

type Point struct{ R, C int }

type IntBoardIterator struct {
	board          *IntBoard
	row, col       int
	rowOffset, pos int
}

func (b *IntBoard) Iterator() IntBoardIterator {
	return IntBoardIterator{board: b}
}
func (b *IntBoard) IteratorAt(row, col int) IntBoardIterator {
	rowOffset := b.Width * row
	return IntBoardIterator{
		board:     b,
		row:       row,
		col:       col,
		rowOffset: rowOffset,
		pos:       rowOffset + col,
	}
}
func (b *IntBoard) IteratorAtPoint(pt Point) IntBoardIterator {
	return b.IteratorAt(pt.R, pt.C)
}

func (i IntBoardIterator) Value() int {
	return i.board.Values[i.pos]
}
func (i IntBoardIterator) Set(value int) {
	i.board.Values[i.pos] = value
}
func (i IntBoardIterator) Point() Point {
	return Point{i.row, i.col}
}
func (i IntBoardIterator) Offset(rows, cols int) IntBoardIterator {
	i.row += rows
	i.rowOffset += i.board.Width * rows
	i.col += cols
	i.pos = i.rowOffset + i.col
	return i
}
func (i IntBoardIterator) NextRC() IntBoardIterator {
	i.pos++
	i.col++
	for i.col >= i.board.Width {
		i.row++
		i.rowOffset += i.board.Width
		i.col -= i.board.Width
	}
	return i
}

func (i IntBoardIterator) SwapBoard(board *IntBoard) IntBoardIterator {
	if i.board.Width != board.Width || i.board.Height != board.Height {
		panic(fmt.Errorf("swap boards requires same size"))
	}
	i.board = board
	return i
}

func (i IntBoardIterator) SquareAdjacencies() []IntBoardIterator {
	if !i.Valid() {
		return nil
	}
	ret := make([]IntBoardIterator, 0, 4)
	if ii := i.Offset(-1, 0); ii.Valid() {
		ret = append(ret, ii)
	}
	if ii := i.Offset(1, 0); ii.Valid() {
		ret = append(ret, ii)
	}
	if ii := i.Offset(0, -1); ii.Valid() {
		ret = append(ret, ii)
	}
	if ii := i.Offset(0, 1); ii.Valid() {
		ret = append(ret, ii)
	}
	return ret
}

func (i IntBoardIterator) DiagonalAdjacencies() []IntBoardIterator {
	if !i.Valid() {
		return nil
	}
	ret := make([]IntBoardIterator, 0, 8)
	for ro := -1; ro <= 1; ro++ {
		for co := -1; co <= 1; co++ {
			if ro == 0 && co == 0 {
				continue
			}
			if ii := i.Offset(ro, co); ii.Valid() {
				ret = append(ret, ii)
			}
		}
	}
	return ret
}

// TODO: NextCR if we need it

func (i IntBoardIterator) Valid() bool {
	return i.row >= 0 && i.col >= 0 &&
		i.row < i.board.Height && i.col < i.board.Width
}
