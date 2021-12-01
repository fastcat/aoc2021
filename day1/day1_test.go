package day1

import (
	_ "embed"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed input.txt
var input string

func Test(t *testing.T) {
	assert.NotEmpty(t, input)
	lines := strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	increases := 0
	last := 0
	for i, l := range lines {
		d, err := strconv.Atoi(l)
		if err != nil {
			panic(err)
		}
		if i > 0 && d > last {
			increases++
		}
		last = d
	}
	t.Log(increases)
}
