package day8

type Analysis struct {
	Patterns       [10]Value
	PatternOptions [10]DigitOption
	WireOptions    [7]Value
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

func (a *Analysis) Rule01BitLengths() {
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

func (a *Analysis) Decode(v Value) DigitOption {
	for i, p := range a.Patterns {
		if v == p {
			return a.PatternOptions[i]
		}
	}
	panic("invalid value")
}
