package day03

import (
	_ "embed"
	"fmt"
	"strings"
	"testing"

	"github.com/fastcat/aoc2021/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

//go:embed example.txt
var exampleInput string

//go:embed input.txt
var input string

type Diagnostic1 struct {
	Gamma, Epsilon, Power int
}

func scan1(t *testing.T, bits int, values, ones []int) Diagnostic1 {
	gamma, mask := 0, 0
	for b := 0; b < bits; b++ {
		mask |= 1 << b
		if len(values)%2 == 0 {
			assert.NotEqual(t, ones[b], len(values)/2, "ambiguous g/e: even split at bit", b)
		}
		if ones[b] > len(values)/2 {
			gamma |= 1 << b
		}
	}
	epsilon := gamma ^ mask
	d := Diagnostic1{
		Gamma:   gamma,
		Epsilon: epsilon,
		Power:   gamma * epsilon,
	}
	t.Logf("scanned %#v", d)
	return d
}

func countInput(t *testing.T, input string) (bits int, values, ones []int) {
	bits, values = parseInput(t, input)
	ones = countOnes(bits, values)
	return
}

func parseInput(t *testing.T, input string) (bits int, values []int) {
	assert.NotEmpty(t, input)
	lines := util.Lines(input)
	require.NotEmpty(t, lines)
	bits = len(lines[0])
	values = util.IntsConv(lines, 2, 32)
	return bits, values
}

func countOnes(bits int, values []int) (ones []int) {
	ones = make([]int, bits)
	for _, v := range values {
		for b := 0; b < bits; b++ {
			if v&(1<<b) != 0 {
				ones[b]++
			}
		}
	}
	return
}

func TestExample1(t *testing.T) {
	bits, values, ones := countInput(t, exampleInput)
	d := scan1(t, bits, values, ones)
	assert.Equal(t, Diagnostic1{22, 9, 198}, d)
}

func TestChallenge1(t *testing.T) {
	bits, values, ones := countInput(t, input)
	scan1(t, bits, values, ones)
}

func binIntStr(values []int, bits int) string {
	binvals := strings.Builder{}
	binvals.WriteString("[]binint{")
	for i, v := range values {
		if i > 0 {
			binvals.WriteString(", ")
		}
		fmtStr := fmt.Sprintf("%%0%db", bits)
		fmt.Fprintf(&binvals, fmtStr, v)
	}
	binvals.WriteString("}")
	return binvals.String()
}

func scan2val(t *testing.T, common bool, bits int, values []int) int {
	// t.Logf("[c:%t] initial %s", common, binIntStr(values, bits))
	// scan bits left to right for filtering
	for b := bits - 1; b >= 0 && len(values) > 1; b-- {
		next := filterByBit(values, b, common)
		require.GreaterOrEqual(t, len(next), 1, "empty?!")
		values = next
		// t.Logf("[c:%t] after bit %d, %s", common, b, binIntStr(values, bits))
	}
	return values[0]
}

func filterByBit(values []int, bit int, common bool) []int {
	ones := 0
	for _, v := range values {
		if v&(1<<bit) != 0 {
			ones++
		}
	}
	ret := make([]int, 0, len(values)/2)
	var want int
	if len(values)%2 == 0 && ones == len(values)/2 {
		// equal split
		if common {
			want = 1
		}
	} else if (ones > len(values)/2) == common {
		want = 1
	}
	want <<= bit
	for _, v := range values {
		if v&(1<<bit) == want {
			ret = append(ret, v)
		}
	}
	return ret
}

type Diagnostic2 struct {
	O2Generator, CO2Scrubber, Rating int
}

func scan2(t *testing.T, bits int, values []int) Diagnostic2 {
	ret := Diagnostic2{
		O2Generator: scan2val(t, true, bits, values),
		CO2Scrubber: scan2val(t, false, bits, values),
	}
	ret.Rating = ret.O2Generator * ret.CO2Scrubber
	t.Logf("analyzed: %#v", ret)
	return ret
}

func TestExample2(t *testing.T) {
	bits, values := parseInput(t, exampleInput)
	d := scan2(t, bits, values)
	assert.Equal(t, Diagnostic2{23, 10, 230}, d)
}

func TestChallenge2(t *testing.T) {
	bits, values := parseInput(t, input)
	scan2(t, bits, values)
}
