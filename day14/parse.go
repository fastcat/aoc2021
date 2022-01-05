package day14

import (
	"fmt"

	"github.com/fastcat/aoc2021/util"
)

type Rule struct {
	Match  [2]rune
	Insert rune
}

func Parse(input string) (seed string, rules []Rule, err error) {
	stanzas := util.Stanzas(input)
	if len(stanzas) != 2 {
		return "", nil, fmt.Errorf("expect 2 stanzas, got %d", len(stanzas))
	}
	seed = stanzas[0]
	lines := util.Lines(stanzas[1])
	rules = make([]Rule, 0, len(lines))
	for _, l := range lines {
		var r Rule
		if _, err = fmt.Sscanf(l, "%c%c -> %c\n", &r.Match[0], &r.Match[1], &r.Insert); err != nil {
			return
		}
		rules = append(rules, r)
	}
	return
}
