package day14

import "strings"

func Apply(seed string, rules []Rule) string {
	b := strings.Builder{}
	b.Grow(len(seed))
	var p rune
	for i, c := range seed {
		if i > 0 {
			for _, r := range rules {
				if r.Match[0] == p && r.Match[1] == c {
					b.WriteRune(r.Insert)
					// assume no duplicate rules
					break
				}
			}
		}
		b.WriteRune(c)
		p = c
	}
	return b.String()
}

func Apply2(counts PairCounts, rules []Rule) {
	// we need to track what we're going to add separately else inserts can affect
	// each other
	inserts := make(PairCounts, len(counts))
	for _, r := range rules {
		c := counts[r.Match]
		counts[r.Match] = 0
		p1, p2 := Pair{r.Match[0], r.Insert}, Pair{r.Insert, r.Match[1]}
		inserts[p1] += c
		inserts[p2] += c
	}
	for p, c := range inserts {
		counts[p] += c
	}
}
