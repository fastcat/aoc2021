package day17

import (
	"math"
)

type Pos struct{ X, Y int }
type Vel struct{ X, Y int }

type Probe struct {
	Pos Pos
	Vel Vel
}

func (p *Probe) Step() Pos {
	p.Pos.X += p.Vel.X
	p.Pos.Y += p.Vel.Y
	if p.Vel.X > 0 {
		p.Vel.X--
	} else if p.Vel.X < 0 {
		p.Vel.X--
	}
	p.Vel.Y--
	return p.Pos
}

type Rect struct {
	X1, X2, Y1, Y2 int
}

func (r Rect) Contains(p Pos) bool {
	return p.X >= r.X1 && p.X <= r.X2 &&
		p.Y >= r.Y1 && p.Y <= r.Y2
}

func (r Rect) Overshot(p Probe) bool {
	return p.Vel.X >= 0 && p.Pos.X > r.X2
}

func (r Rect) Undershot(p Probe) bool {
	return p.Vel.X <= 0 && p.Pos.X < r.X1
}

func (r Rect) Sank(p Probe) bool {
	return p.Vel.Y <= 0 && p.Pos.Y < r.Y1
}

func (p Probe) WillHit(target Rect) (bool, int, Pos) {
	// this is a value receiver so we work on a copy
	maxY := p.Pos.Y
	for {
		// fmt.Println("probe at", p.Pos, "going", p.Vel)
		if target.Contains(p.Pos) {
			return true, maxY, p.Pos
		}
		if target.Overshot(p) {
			return false, maxY, p.Pos
		}
		if target.Undershot(p) {
			return false, maxY, p.Pos
		}
		if target.Sank(p) {
			return false, maxY, p.Pos
		}
		p.Step()
		if p.Pos.Y > maxY {
			maxY = p.Pos.Y
		}
	}
}

func (p Probe) OptimalVelocityRange(target Rect) (fast, slow Vel) {
	// sum(1..n)=n(n+1)/2

	// after lofting up, we'll return to the initial y pos but with inverted
	// velocity plus 1, which should then drop us to the bottom of the target
	if target.Y1 >= p.Pos.Y {
		panic(nil)
	}
	y1 := p.Pos.Y - target.Y1 - 1
	y2 := p.Pos.Y - target.Y2 - 1

	// try to hit the right edge of the target, round down
	dx1 := float64(target.X2 - p.Pos.X)
	dx2 := float64(target.X1 - p.Pos.X)
	// n^2 + n - 2dx = 0
	xf1 := math.Sqrt(1+8*dx1)/2 - 0.5
	xf2 := math.Sqrt(1+8*dx2)/2 - 0.5
	x1 := int(math.Floor(xf1))
	x2 := int(math.Ceil(xf2))
	_ = x1

	return Vel{x1, y1}, Vel{x2, y2}
}
