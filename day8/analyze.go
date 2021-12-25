package day8

import (
	"fmt"
	"strings"
)

type Analysis struct {
	Patterns       [10]Value
	PatternOptions [10]DigitOption
	WireOptions    [7]Value
}

func (a *Analysis) String() string {
	b := &strings.Builder{}
	for i := 0; i < 7; i++ {
		fmt.Fprintf(b, "%c %s\n", 'a'+rune(i), strings.ToUpper(a.WireOptions[i].String()))
	}
	for i := 0; i < 10; i++ {
		fmt.Fprintf(b, "%s\t%s\n", a.Patterns[i], a.PatternOptions[i])
	}
	return b.String()
}

func NewAnalysis(e Entry) Analysis {
	var ret Analysis
	ret.Patterns = e.Patterns
	for i := range ret.PatternOptions {
		ret.PatternOptions[i] = AllDigits
	}
	for i := range ret.WireOptions {
		ret.WireOptions[i] = AllValues
	}
	return ret
}

func (a *Analysis) Done() bool {
	for _, o := range a.PatternOptions {
		if o.Len() != 1 {
			return false
		}
	}
	for _, o := range a.WireOptions {
		if o.Len() != 1 {
			return false
		}
	}
	return true
}

func (a *Analysis) maskWires(pattern, mask Value) {
	for i := 0; i < 7; i++ {
		v := Value(1 << i)
		if pattern&v == 0 {
			continue
		}
		a.WireOptions[i] &= mask
	}
}

func (a *Analysis) Rule01DigitOptions() {
	for i, p := range a.Patterns {
		// TODO: apply wire option masks
		switch p.Len() {
		case 2:
			a.PatternOptions[i] &= One
		case 3:
			a.PatternOptions[i] &= Seven
		case 4:
			a.PatternOptions[i] &= Four
		case 5:
			a.PatternOptions[i] &= Two | Three | Five
		case 6:
			a.PatternOptions[i] &= Zero | Six | Nine
		case 7:
			a.PatternOptions[i] &= Eight
		default:
			panic("invalid observation")
		}
	}
}

func (a *Analysis) Rule02WireMasks() {
	for _, p := range a.Patterns {
		// TODO: apply wire option masks
		switch p.Len() {
		case 2:
			a.maskWires(p, C|F)
			a.maskWires(^p, ^(C | F))
		case 3:
			a.maskWires(p, A|C|F)
			a.maskWires(^p, ^(A | C | F))
		case 4:
			a.maskWires(p, B|C|D|F)
			a.maskWires(^p, ^(B | C | D | F))
		case 5:
		case 6:
			a.maskWires(^p, C|D|E)
		case 7:
		default:
			panic("invalid observation")
		}
	}
}

func (a *Analysis) Analyze() bool {
	a.Rule01DigitOptions()
	a.Rule02WireMasks()
	return a.Done()
}

func (a *Analysis) Decode(values ...Value) []DigitOption {
	ret := make([]DigitOption, 0, len(values))
VALUES:
	for _, v := range values {
		for i, p := range a.Patterns {
			if v == p {
				ret = append(ret, a.PatternOptions[i])
				continue VALUES
			}
		}
		panic("invalid value")
	}
	return ret
}
