package day15

import (
	"fmt"
	"math"

	"github.com/fastcat/aoc2021/util"
)

type PathCache struct {
	entryCosts  *util.IntBoard
	pathCosts   map[util.Point]int
	destination util.Point
}

func PathTo(entryCosts *util.IntBoard, dest util.Point) *PathCache {
	return &PathCache{entryCosts, make(map[util.Point]int, len(entryCosts.Values)), dest}
}

func (p *PathCache) Valid() bool {
	return p != nil &&
		p.entryCosts.Valid() &&
		p.pathCosts != nil &&
		p.entryCosts.IteratorAtPoint(p.destination).Valid()
}

func (p *PathCache) Fill() {
	// seed: the destination point always has cost 0
	p.pathCosts[p.destination] = 0
	// iterate improving the cost map until we stop finding any improvements. at
	// each iteration, we only need to look at the points adjacent to those we
	// improved on the prior iteration.
	for improved := (map[util.Point]bool{p.destination: true}); len(improved) != 0; {
		nextImproved := make(map[util.Point]bool)
		for pt := range improved {
			i := p.entryCosts.IteratorAtPoint(pt)
			// iv is the current best cost for going from pt(i) to the dest
			iv := p.pathCosts[pt]
			for _, j := range i.SquareAdjacencies() {
				if jv, ok := p.pathCosts[j.Point()]; ok {
					// we have a known cost jv to the dest for point j adjacent to i. see
					// if going from i to j and from there to the dest is better than the
					// current best path for i to the dest.
					// ev is the cost to enter point j adjacent to i
					ev := j.Value()
					if jv+ev < iv {
						p.pathCosts[pt] = jv + ev
						nextImproved[pt] = true
					}
				} else {
					// we don't know a cost for going from j to the dest, so start out
					// with going from j to i to the dest
					jp := j.Point()
					p.pathCosts[jp] = i.Value() + iv
					nextImproved[jp] = true
				}
			}
			improved = nextImproved
		}
	}
}

func (p *PathCache) fillSlow() *util.IntBoard {
	pathCosts := &util.IntBoard{
		Values: make([]int, len(p.entryCosts.Values)),
		Width:  p.entryCosts.Width,
		Height: p.entryCosts.Height,
	}
	for i := pathCosts.Iterator(); i.Valid(); i = i.NextRC() {
		i.Set(math.MaxInt)
	}
	// seed: the destination point always has cost 0
	pathCosts.IteratorAtPoint(p.destination).Set(0)
	// iterate improving the cost map until we stop finding any improvements
	for improved := true; improved; {
		improved = false
		for i := pathCosts.Iterator(); i.Valid(); i = i.NextRC() {
			iv := i.Value()
			for _, j := range i.SquareAdjacencies() {
				jv := j.Value()
				if jv < math.MaxInt {
					ev := j.SwapBoard(p.entryCosts).Value()
					if jv+ev < iv {
						i.Set(jv + ev)
						improved = true
					}
				}
			}
		}
	}
	for i := pathCosts.Iterator(); i.Valid(); i = i.NextRC() {
		p.pathCosts[i.Point()] = i.Value()
	}
	return pathCosts
}

func (p *PathCache) costMap() *util.IntBoard {
	pathCosts := &util.IntBoard{
		Values: make([]int, len(p.entryCosts.Values)),
		Width:  p.entryCosts.Width,
		Height: p.entryCosts.Height,
	}
	for i := pathCosts.Iterator(); i.Valid(); i = i.NextRC() {
		i.Set(math.MaxInt)
	}
	for pt, v := range p.pathCosts {
		pathCosts.IteratorAtPoint(pt).Set(v)
	}
	return pathCosts
}

func (p *PathCache) CostFrom(start util.Point) int {
	if v, ok := p.pathCosts[start]; ok {
		return v
	}
	panic(fmt.Errorf("invalid start point %v", start))
}
