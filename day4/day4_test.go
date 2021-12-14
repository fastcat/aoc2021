package day4

import (
	_ "embed"
	"fmt"
	"strings"
	"testing"

	"github.com/fastcat/aoc2021/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

//go:embed example.txt
var exampleInput string

//go:embed input.txt
var input string

func TestData(t *testing.T) {
	assert.NotEmpty(t, exampleInput, "example input")
	assert.NotEmpty(t, input, "challenge input")
}

// [row][col]val
type board = [5][5]int

func parseBoard(t *testing.T, in string) board {
	var ret board
	// lol
	n, err := fmt.Sscanf(
		in,
		"%d %d %d %d %d\n"+
			"%d %d %d %d %d\n"+
			"%d %d %d %d %d\n"+
			"%d %d %d %d %d\n"+
			"%d %d %d %d %d\n",
		&ret[0][0], &ret[0][1], &ret[0][2], &ret[0][3], &ret[0][4],
		&ret[1][0], &ret[1][1], &ret[1][2], &ret[1][3], &ret[1][4],
		&ret[2][0], &ret[2][1], &ret[2][2], &ret[2][3], &ret[2][4],
		&ret[3][0], &ret[3][1], &ret[3][2], &ret[3][3], &ret[3][4],
		&ret[4][0], &ret[4][1], &ret[4][2], &ret[4][3], &ret[4][4],
	)
	require.NoError(t, err)
	require.Equal(t, 25, n)
	return ret
}

func TestParseBoard(t *testing.T) {
	assert.Equal(t,
		board{
			{1, 2, 3, 4, 5},
			{6, 7, 8, 9, 10},
			{11, 12, 13, 14, 15},
			{16, 17, 18, 19, 20},
			{21, 22, 23, 24, 25},
		},
		parseBoard(t,
			"1 2 3 4 5\n"+
				"6 7 8 9 10\n"+
				"11 12 13 14 15\n"+
				"16 17 18 19 20\n"+
				"21 22 23 24 25\n",
		),
	)
}

type Input struct {
	Calls  []int
	Boards []board
}

func parseInput(t *testing.T, input string) Input {
	assert.NotEmpty(t, input)
	stanzas := strings.Split(input, "\n\n")
	assert.GreaterOrEqual(t, len(stanzas), 2)
	callStrs := strings.Split(stanzas[0], ",")
	var ret Input
	ret.Calls = util.Ints(callStrs)
	ret.Boards = make([]board, len(stanzas)-1)
	for i, s := range stanzas[1:] {
		ret.Boards[i] = parseBoard(t, s)
	}
	return ret
}

type BoardResult struct {
	WinsAfter int // 0 = never wins
	Score     int // only valid if it wins
}

func hasWon(b board, called map[int]bool) bool {
	// check each row
	for _, row := range b {
		won := true
		for _, cell := range row {
			if !called[cell] {
				won = false
				break
			}
		}
		if won {
			return true
		}
	}
	// check each column
	for col := 0; col < len(b[0]); col++ {
		won := true
		for _, row := range b {
			if !called[row[col]] {
				won = false
				break
			}
		}
		if won {
			return true
		}
	}
	return false
}

func TestHasWon(t *testing.T) {
	tests := []struct {
		name   string
		board  board
		called []int
		want   bool
	}{
		{
			"example not winner yet",
			parseBoard(t, "14 21 17 24  4\n10 16 15  9 19\n18  8 23 26 20\n22 11 13  6  5\n 2  0 12  3  7\n"),
			[]int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21},
			false,
		},
		{
			"example winner",
			parseBoard(t, "14 21 17 24  4\n10 16 15  9 19\n18  8 23 26 20\n22 11 13  6  5\n 2  0 12  3  7\n"),
			[]int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			called := make(map[int]bool, len(tt.called))
			for _, c := range tt.called {
				called[c] = true
			}
			got := hasWon(tt.board, called)
			assert.Equal(t, tt.want, got)
		})
	}
}

func scoreBoard(b board, unmarked map[int]bool, call int) int {
	score := 0
	for _, row := range b {
		for _, cell := range row {
			if unmarked[cell] {
				score += cell
			}
		}
	}
	return score * call
}

func playBoard(b board, calls []int) BoardResult {
	unmarked := make(map[int]bool, 25)
	for _, row := range b {
		for _, cell := range row {
			unmarked[cell] = true
		}
	}
	called := make(map[int]bool, 25)
	for i, call := range calls {
		called[call] = true
		unmarked[call] = false
		if hasWon(b, called) {
			return BoardResult{
				WinsAfter: i + 1,
				Score:     scoreBoard(b, unmarked, call),
			}
		}
	}
	return BoardResult{}
}

func TestExample1(t *testing.T) {
	firstWinner, firstWinsAfter, firstWinsScore := fastestWinner(t, exampleInput)
	assert.Equal(t, 2, firstWinner)
	assert.Equal(t, 12, firstWinsAfter)
	assert.Equal(t, 4512, firstWinsScore)
}

func fastestWinner(t *testing.T, input string) (firstWinner, firstWinsAfter, firstWinsScore int) {
	in := parseInput(t, input)
	firstWinner = -1
	firstWinsAfter = -1
	firstWinsScore = -1
	for i, b := range in.Boards {
		r := playBoard(b, in.Calls)
		if r.WinsAfter == 0 {
			t.Logf("board %d is a loser", i)
			continue
		}
		if firstWinsAfter < 0 || r.WinsAfter < firstWinsAfter {
			t.Logf("board %d wins faster than %d: %d < %d, score %d vs %d",
				i, firstWinner,
				r.WinsAfter, firstWinsAfter,
				r.Score, firstWinsScore,
			)
			firstWinner = i
			firstWinsAfter = r.WinsAfter
			firstWinsScore = r.Score
		}
	}
	if assert.GreaterOrEqual(t, firstWinner, 0) {
		t.Logf("best board is %d after %d scoring %d", firstWinner, firstWinsAfter, firstWinsScore)
	}
	return
}

func TestChallenge1(t *testing.T) {
	fastestWinner(t, input)
}
