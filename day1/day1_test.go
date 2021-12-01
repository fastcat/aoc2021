package day1

import (
	_ "embed"
	"strconv"
	"testing"

	"github.com/fastcat/aoc2021/util"
	"github.com/stretchr/testify/assert"
)

//go:embed input.txt
var input string

func TestMain(t *testing.T) {
	assert.NotEmpty(t, input)
	lines := util.Lines(input)
	ints := util.Ints(lines)
	increases := incsP1(ints)
	t.Log("Part 1:", increases)
	assert.Equal(t, increases, 1288)

	increases = incsP2(ints)
	t.Log("Part 2:", increases)
}

func incsP2(ints []int) int {
	lastsum := 0
	increases := 0
	for i, d := range ints {
		if i < 3 {
			lastsum += d
			continue
		}
		nextsum := lastsum - ints[i-3] + ints[i]
		if nextsum > lastsum {
			increases++
		}
		lastsum = nextsum
	}
	return increases
}

func TestIncsP2(t *testing.T) {
	tests := []struct {
		arg  []int
		want int
	}{
		{[]int{1, 2, 3, 4}, 1},
		{[]int{1, 2, 3, 1}, 0},
		{[]int{607, 618, 618, 617, 647, 716, 769, 792}, 5},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			assert.Equal(t, tt.want, incsP2(tt.arg))
		})
	}
}

func incsP1(ints []int) int {
	increases := 0
	last := 0
	for i, d := range ints {
		if i > 0 && d > last {
			increases++
		}
		last = d
	}
	return increases
}

func TestIncsP1(t *testing.T) {
	assert.Equal(t, incsP1([]int{1, 2, 3}), 2)
	assert.Equal(t, incsP1([]int{1, 3, 2, 4}), 2)
}
