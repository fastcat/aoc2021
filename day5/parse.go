package day5

import (
	"fmt"

	"github.com/fastcat/aoc2021/util"
)

type Point struct{ X, Y int }
type Line struct{ Start, End Point }

func ParseLines(input string) ([]Line, error) {
	inputLines := util.Lines(input)
	parsedLines := make([]Line, 0, len(inputLines))
	for _, l := range inputLines {
		var pl Line
		if n, err := fmt.Sscanf(l, "%d,%d -> %d,%d",
			&pl.Start.X, &pl.Start.Y,
			&pl.End.X, &pl.End.Y,
		); err != nil {
			return parsedLines, err
		} else if n != 4 {
			return parsedLines, fmt.Errorf("corrupt line")
		} else {
			parsedLines = append(parsedLines, pl)
		}
	}
	return parsedLines, nil
}
