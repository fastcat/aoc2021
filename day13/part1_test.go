package day13

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1Example(t *testing.T) {
	board, folds, err := ParseBoardAndFolds(exampleInput)
	assert.NoError(t, err)
	t.Logf("initial\n%s", board)
	for _, fold := range folds {
		board.ApplyFold(fold)
		t.Logf("after %s\n%s", fold, board)
	}
	assert.Equal(t, "#####\n#...#\n#...#\n#...#\n#####", board.String())
}

func TestPart1Challenge(t *testing.T) {
	board, folds, err := ParseBoardAndFolds(challengeInput)
	assert.NoError(t, err)
	// t.Logf("initial\n%s", board)
	board.ApplyFold(folds[0])
	// t.Logf("after one fold\n%s", board)
	t.Logf("challenge cells set: %d", len(board.cells))
}

func TestPart2Challenge(t *testing.T) {
	board, folds, err := ParseBoardAndFolds(challengeInput)
	assert.NoError(t, err)
	board.ApplyFolds(folds)
	t.Logf("after final\n%s", board)
}
