package day9

import (
	_ "embed"
	"fmt"
	"sort"

	"github.com/fastcat/aoc2021/util"
)

//go:embed example.txt
var exampleInput string

//go:embed input.txt
var challengeInput string

type Board struct {
	width, height int
	depths        []int
}

func ParseBoard(input string) (*Board, error) {
	rows := util.Lines(input)
	board := &Board{height: len(rows)}
	for r, row := range rows {
		if board.width != 0 {
			if board.width != len(row) {
				return board, fmt.Errorf("row %d wrong width %d should be %d", r, len(row), board.width)
			}
		} else {
			board.width = len(row)
			board.depths = make([]int, board.width*board.height)
		}
		ro := board.width * r
		for c := 0; c < len(row); c++ {
			depth := int(row[c] - '0')
			if depth < 0 || depth > 9 {
				return board, fmt.Errorf("r,c %d,%d bad value %c=%d", r, c, row[c], depth)
			}
			board.depths[ro+c] = depth
		}
	}
	return board, nil
}

type Point struct{ R, C int }

func (b *Board) LocalMinima() []Point {
	var minima []Point
	for r, ro := 0, 0; r < b.height; r, ro = r+1, ro+b.width {
		for c, p := 0, ro; c < b.width; c, p = c+1, p+1 {
			if c > 0 && b.depths[p] >= b.depths[p-1] {
				// fmt.Printf("%d,%d nlm left\n", r, c)
				continue // not less than left
			}
			if c < b.width-1 && b.depths[p] >= b.depths[p+1] {
				// fmt.Printf("%d,%d nlm right\n", r, c)
				continue // not less than right
			}
			if r > 0 && b.depths[p] >= b.depths[p-b.width] {
				// fmt.Printf("%d,%d nlm above\n", r, c)
				continue // not less than above
			}
			if r < b.height-1 && b.depths[p] >= b.depths[p+b.width] {
				// fmt.Printf("%d,%d nlm below\n", r, c)
				continue // not less than below
			}
			minima = append(minima, Point{r, c})
		}
	}
	return minima
}

func (b *Board) DepthAt(p Point) int {
	if p.R < 0 || p.R >= b.height || p.C < 0 || p.C >= b.width {
		panic(fmt.Errorf("invalid point %v", p))
	}
	return b.depths[p.C+b.width*p.R]
}

func Part1LocalMinimaRiskLevelSum(b *Board) int {
	m := b.LocalMinima()
	sum := 0
	for _, p := range m {
		sum += b.DepthAt(p) + 1
	}
	return sum
}

type Basin struct {
	Minimum Point
	Size    int
}

func (b *Board) Basins() []Basin {
	// for each point on the board, walk the gradient descent until we hit a local
	// minimum, and then mark everything on that path as a part of a basin with
	// that minimum. If, in that walk, we hit a point with an already known basin,
	// then we can stop there and mark our partial path as also part of that
	// basin.
	basins := map[Point]int{}
	paths := make([]*Point, len(b.depths))
	for r, ro := 0, 0; r < b.height; r, ro = r+1, ro+b.width {
		for c, p := 0, ro; c < b.width; c, p = c+1, p+1 {
			// skip if we we already walked this point
			if paths[p] != nil {
				continue
			}
			// skip high points
			if b.depths[p] >= 9 {
				continue
			}
			pathLen := 0
			minimum := Point{-1, -1}
			for pr, pc, pp := r, c, p; ; {
				pathLen++
				// record everything on the path as part of whatever minimum we
				// eventually find
				paths[pp] = &minimum
				nextpr, nextpc, nextpp := pr, pc, pp
				if pc > 0 && b.depths[pp-1] < b.depths[pp] {
					// left < cur
					nextpp = pp - 1
					nextpc = pc - 1
				}
				if c < b.width-1 && b.depths[pp+1] < b.depths[nextpp] {
					// right < best
					nextpp = pp + 1
					nextpc = pc + 1
				}
				if pr > 0 && b.depths[pp-b.width] < b.depths[nextpp] {
					// above < best, might need to reset C
					nextpp = pp - b.width
					nextpc = pc
					nextpr = pr - 1
				}
				if pr < b.height-1 && b.depths[pp+b.width] < b.depths[nextpp] {
					nextpp = pp + b.width
					nextpc = pc
					nextpr = pr + 1
				}
				if nextpp == pp {
					// hit a local minimum
					minimum = Point{pr, pc}
					break
				} else if m := paths[nextpp]; m != nil {
					// hit a prior path
					minimum = *m
					break
				}
				pr, pc, pp = nextpr, nextpc, nextpp
			}
			// path now contains all the new points, and minimum points to the basin
			// minimum.
			basins[minimum] += pathLen
		}
	}
	ret := make([]Basin, 0, len(basins))
	for m, s := range basins {
		ret = append(ret, Basin{m, s})
	}
	return ret
}

func Part2BasinScore(basins []Basin) int {
	// sort the basins descending by size
	sort.Slice(basins, func(i, j int) bool { return basins[i].Size > basins[j].Size })
	product := 1
	for i := 0; i < 3 && i < len(basins); i++ {
		product *= basins[i].Size
	}
	return product
}
