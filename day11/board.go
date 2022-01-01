package day11

import (
	"fmt"

	"github.com/fastcat/aoc2021/util"
)

type Board struct {
	levels        []int
	width, height int
}

func Parse(input string) (Board, error) {
	lines := util.Lines(input)
	board := Board{
		width:  len(lines[0]),
		height: len(lines),
	}
	board.levels = make([]int, board.width*board.height)
	ro := 0
	for r, l := range lines {
		for c, v := range l {
			if v < '0' || v > '9' {
				return board, fmt.Errorf("invalid cell %d,%d=%c", r, c, v)
			}
			board.levels[ro+c] = int(v - '0')
		}
		ro += board.width
	}
	return board, nil
}

func MustParse(input string) Board {
	board, err := Parse(input)
	if err != nil {
		panic(err)
	}
	return board
}

func (b *Board) Step() (flashes int) {
	panic("wip")
}
