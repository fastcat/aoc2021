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
			// iv is the current best cost for i->dest, which was just improved
			iv := p.pathCosts[pt]
			// ivv is the cost for N->i->dest, for some neighbor N
			ivv := i.Value() + iv
			// having just improved i->dest might improve its neighbors
			for _, j := range i.SquareAdjacencies() {
				jp := j.Point()
				if jv, ok := p.pathCosts[jp]; ok {
					// jv is the cost for j->dest.
					// see if i->j->dest is better than the current i->dest. technically
					// we could skip this, since we're only here because the i->dest path
					// just improved, but adding this check is a big performance
					// improvement.
					if jvv := j.Value() + jv; jvv < iv {
						p.pathCosts[pt] = jvv
						nextImproved[pt] = true
					}
					// see if j->i->dest is better than the current j->dest
					if ivv < jv {
						p.pathCosts[jp] = ivv
						nextImproved[jp] = true
					}
				} else {
					// we don't know a cost for j->dest, so seed it as j->i->dest
					p.pathCosts[jp] = ivv
					nextImproved[jp] = true
				}
			}
		}
		improved = nextImproved
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
