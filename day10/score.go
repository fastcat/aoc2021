package day10

import (
	"fmt"
	"sort"
)

func Part1Score(input string, err error) int {
	isInv, pe := IsInvalid(err)
	if !isInv {
		return 0
	}
	c := input[pe.Position().Offset]
	switch c {
	case ')':
		return 3
	case ']':
		return 57
	case '}':
		return 1197
	case '>':
		return 25137
	default:
		panic(fmt.Errorf("no score for '%c'", c))
	}
}

func Part2ScoreInc(score int, open string) int {
	score *= 5
	switch open {
	case "(":
		return score + 1
	case "[":
		return score + 2
	case "{":
		return score + 3
	case "<":
		return score + 4
	default:
		panic(fmt.Errorf("invalud open '%s'", open))
	}
}

func Part2LineScore(line string) (bool, int) {
	var parsed ChunkList
	err := chunksParser.ParseString("", line, &parsed)
	ok, _ := IsIncomplete(err)
	if !ok {
		return false, 0
	}
	// scoring goes from inside out, so make a stack and then score it
	var incomplete []Chunk
	for c := parsed.Contents[len(parsed.Contents)-1]; c.Close == ""; c = c.Contents[len(c.Contents)-1] {
		incomplete = append(incomplete, c)
		if len(c.Contents) == 0 {
			break
		}
	}
	score := 0
	for i := len(incomplete) - 1; i >= 0; i-- {
		score = Part2ScoreInc(score, incomplete[i].Open)
	}
	return true, score
}

func Part2TotalScore(lines ...string) int {
	scores := make([]int, 0, len(lines))
	for _, line := range lines {
		if ok, score := Part2LineScore(line); ok {
			scores = append(scores, score)
		}
	}
	sort.Ints(scores)
	mp := len(scores) / 2
	return scores[mp]
}
