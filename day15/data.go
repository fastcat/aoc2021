package day15

import (
	_ "embed"

	"github.com/fastcat/aoc2021/util"
)

//go:embed example.txt
var exampleInput string

//go:embed example5.txt
var exampleFiveXInput string

//go:embed challenge.txt
var challengeInput string

func BuildFiveX(oneX *util.IntBoard) *util.IntBoard {
	fiveX := &util.IntBoard{
		Values: make([]int, len(oneX.Values)*25),
		Width:  oneX.Width * 5,
		Height: oneX.Height * 5,
	}
	i5 := make([]struct {
		po util.Point
		vo int
	}, 0, 25)
	for rp := 0; rp < 5; rp++ {
		for cp := 0; cp < 5; cp++ {
			i5 = append(i5, struct {
				po util.Point
				vo int
			}{util.Point{R: rp * oneX.Height, C: cp * oneX.Width}, rp + cp})
		}
	}
	for i := oneX.Iterator(); i.Valid(); i = i.NextRC() {
		v := i.Value()
		ip := i.Point()
		for _, j := range i5 {
			jv := v + j.vo
			jv = (jv-1)%9 + 1
			fiveX.IteratorAtPoint(util.Point{R: ip.R + j.po.R, C: ip.C + j.po.C}).Set(jv)
		}
	}
	return fiveX
}
