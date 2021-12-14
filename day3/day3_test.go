package day3

import (
	_ "embed"
	"testing"

	"github.com/fastcat/aoc2021/util"
	"github.com/stretchr/testify/assert"
)

//go:embed example.txt
var exampleInput string

//go:embed input.txt
var input string

type Diagnostic struct {
	Gamma, Epsilon, Power int
}

func scan(t *testing.T, input string) Diagnostic {
	lines := util.Lines(input)
	bits := len(lines[0])
	values := util.IntsConv(lines, 2, 32)
	ones := make([]int, bits)
	for _, v := range values {
		for b := 0; b < bits; b++ {
			if v&(1<<b) != 0 {
				ones[b]++
			}
		}
	}
	gamma, mask := 0, 0
	for b := 0; b < bits; b++ {
		mask |= 1 << b
		if ones[b] > len(values)/2 {
			gamma |= 1 << b
		}
	}
	epsilon := gamma ^ mask
	d := Diagnostic{
		Gamma:   gamma,
		Epsilon: epsilon,
		Power:   gamma * epsilon,
	}
	t.Logf("scanned %#v", d)
	return d
}

func TestExample1(t *testing.T) {
	assert.NotEmpty(t, exampleInput)
	d := scan(t, exampleInput)
	assert.Equal(t, Diagnostic{22, 9, 198}, d)
}

func TestChallenge1(t *testing.T) {
	assert.NotEmpty(t, input)
	scan(t, input)
}
