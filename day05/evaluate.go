package day05

import (
	"fmt"
	"math"
)

func GridOverlaps(lines []Line, includeDiagonal bool) map[Point]int {
	grid := make(map[Point]int)
	for _, l := range lines {
		updateGrid(grid, l, includeDiagonal)
	}
	return grid
}

func updateGrid(grid map[Point]int, line Line, includeDiagnoal bool) {
	if line.Start.X == line.End.X {
		start, end := line.Start.Y, line.End.Y
		if end < start {
			start, end = end, start
		}
		for p := (Point{line.Start.X, start}); p.Y <= end; p.Y++ {
			grid[p] = grid[p] + 1
		}
	} else if line.Start.Y == line.End.Y {
		start, end := line.Start.X, line.End.X
		if end < start {
			start, end = end, start
		}
		for p := (Point{start, line.Start.Y}); p.X <= end; p.X++ {
			grid[p] = grid[p] + 1
		}
	} else if math.Abs(float64(line.End.X-line.Start.X)) == math.Abs(float64(line.End.Y-line.Start.Y)) {
		if includeDiagnoal {
			start, end := line.Start, line.End
			if start.X > end.X {
				start, end = end, start
			}
			// x is now increasing, but y may be decreasing
			yDelta := 1
			if start.Y > end.Y {
				yDelta = -1
			}
			for p := start; p.X <= end.X; p.X, p.Y = p.X+1, p.Y+yDelta {
				grid[p] = grid[p] + 1
			}
		}
	} else {
		panic(fmt.Errorf("wonky diagonal line: %v", line))
	}
}

func NumPointsAtLeast(grid map[Point]int, threshold int, log func(string, ...interface{})) int {
	nAbove := 0
	for p, hits := range grid {
		if hits >= threshold {
			if log != nil {
				log("point %v has %d hits", p, hits)
			}
			nAbove++
		}
	}
	return nAbove
}
