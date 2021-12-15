package day5

func GridOverlaps(lines []Line) map[Point]int {
	grid := make(map[Point]int)
	for _, l := range lines {
		updateGrid(grid, l)
	}
	return grid
}

func updateGrid(grid map[Point]int, line Line) {
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
	} else {
		// panic(fmt.Errorf("diagonal line: %v", line))
		// continue
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
