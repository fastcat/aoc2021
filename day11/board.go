package day11

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fastcat/aoc2021/util"
)

type Board struct {
	levels        []int
	width, height int
}

func (b *Board) String() string {
	var sb strings.Builder
	for i := 0; i < len(b.levels); i++ {
		if i > 0 && i%b.width == 0 {
			sb.WriteRune('\n')
		}
		sb.WriteString(strconv.Itoa(b.levels[i]))
	}
	return sb.String()
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
	b.Increment()
	return b.Flashes()
}

func (b *Board) Increment() {
	for i := range b.levels {
		b.levels[i]++
	}
}

func (b *Board) Flashes() int {
	flashes := 0
	for done := false; !done; {
		done = true
		for r, ro := 0, 0; r < b.height; r, ro = r+1, ro+b.width {
			for c, p := 0, ro; c < b.width; c, p = c+1, p+1 {
				if b.levels[p] <= 9 {
					continue
				}
				flashes++
				cont := b.flash(r, ro, c, p)
				if cont {
					done = false
				}
			}
		}
	}
	return flashes
}
func (b *Board) flash(r, ro, c, p int) (cont bool) {
	b.levels[p] = 0
	for dro := -b.width; dro <= b.width; dro += b.width {
		if ro+dro < 0 || ro+dro >= len(b.levels) {
			continue
		}
		for dc, pp := -1, p+dro-1; dc <= 1; dc, pp = dc+1, pp+1 {
			if dro == 0 && pp == p {
				continue
			} else if c+dc < 0 || c+dc >= b.width {
				continue
			}
			if b.levels[pp] != 0 {
				b.levels[pp]++
				if b.levels[pp] > 9 {
					cont = true
				}
			}
		}
	}
	return
}

func (b *Board) AllZero() bool {
	for _, l := range b.levels {
		if l > 0 {
			return false
		}
	}
	return true
}
