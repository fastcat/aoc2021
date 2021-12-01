package util

import (
	"strconv"
	"strings"
)

func Lines(input string) []string {
	return strings.Split(strings.TrimSuffix(input, "\n"), "\n")
}

func Ints(lines []string) []int {
	ret := make([]int, len(lines))
	for i, l := range lines {
		j, err := strconv.Atoi(l)
		if err != nil {
			panic(err)
		}
		ret[i] = j
	}
	return ret
}
