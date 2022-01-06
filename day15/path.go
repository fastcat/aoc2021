package day15

import (
	"fmt"
	"math"

	"github.com/fastcat/aoc2021/util"
)

type PathCache struct {
	entryCosts  *util.IntBoard
	pathCosts   *util.IntBoard
	destination util.Point
}

func PathTo(entryCosts *util.IntBoard, dest util.Point) *PathCache {
	pathCosts := &util.IntBoard{
		Values: make([]int, len(entryCosts.Values)),
		Width:  entryCosts.Width,
		Height: entryCosts.Height,
	}
	di := pathCosts.IteratorAtPoint(dest)
	di.Set(di.SwapBoard(entryCosts).Value())
	return &PathCache{entryCosts, pathCosts, dest}
}

func (p *PathCache) Valid() bool {
	return p != nil &&
		p.entryCosts.Valid() &&
		p.pathCosts.Valid() &&
		p.entryCosts.IteratorAtPoint(p.destination).Valid()
}

func (p *PathCache) Fill() {
	for i := p.pathCosts.Iterator(); i.Valid(); i = i.NextRC() {
		i.Set(math.MaxInt)
	}
	// seed: the destination point always has cost 0
	p.pathCosts.IteratorAtPoint(p.destination).Set(0)
	// iterate improving the cost map until we stop finding any improvements
	for improved := true; improved; {
		improved = false
		for i := p.pathCosts.Iterator(); i.Valid(); i = i.NextRC() {
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
}

func (p *PathCache) CostFrom(start util.Point) int {
	i := p.pathCosts.IteratorAtPoint(start)
	if !i.Valid() {
		panic(fmt.Errorf("invalid start point %v", start))
	}
	return i.Value()
}
