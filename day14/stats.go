package day14

import "sort"

type Stat struct {
	Element rune
	Count   int64
}

func Analyze(polymer string) []Stat {
	accum := map[rune]int64{}
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

func Analyze2(seed string, counts PairCounts) []Stat {
	accum := map[rune]int64{}
	for e, c := range counts {
		accum[e[0]] += c
		accum[e[1]] += c
	}
	// everything except the first & last element are double-counted. the first
	// and last are double counted less 1
	first, last := rune(seed[0]), rune(seed[len(seed)-1])
	for e, c := range accum {
		if e == first || e == last {
			if first == last {
				// actually keep two extra
				accum[e] -= (c - 2) / 2
			} else {
				accum[e] -= (c - 1) / 2
			}
		} else {
			accum[e] = c / 2
		}
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
