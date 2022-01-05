package day14

import "sort"

type Stat struct {
	Element rune
	Count   int
}

func Analyze(polymer string) []Stat {
	accum := map[rune]int{}
	for _, e := range polymer {
		accum[e]++
	}
	stats := make([]Stat, 0, len(accum))
	for e, c := range accum {
		stats = append(stats, Stat{e, c})
	}
	sort.Slice(stats, func(i, j int) bool {
		return stats[i].Count < stats[j].Count
	})
	return stats
}
