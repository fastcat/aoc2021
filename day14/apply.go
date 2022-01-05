package day14

import "strings"

func Apply(seed string, rules ...Rule) string {
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
