package day8

import "math/bits"

type Value uint8

const (
	A Value = 1 << iota
	B
	C
	D
	E
	F
	G
	AllValues Value = 0b1111111
)

type Entry struct {
	Patterns [10]Value
	Outputs  [4]Value
}

func (v Value) Len() int {
	return bits.OnesCount8(uint8(v))
}

type DigitOption uint16

func (d DigitOption) Len() int {
	return bits.OnesCount16(uint16(d))
}

const (
	Zero DigitOption = 1 << iota
	One
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	AllDigits DigitOption = 0b1111111111
)

type DigitDisplay Value

const (
	DisplayZero  = DigitDisplay(A | B | C | E | F | G)
	DisplayOne   = DigitDisplay(C | F)
	DisplayTwo   = DigitDisplay(A | C | D | E | G)
	DisplayThree = DigitDisplay(A | C | D | F | G)
	DisplayFour  = DigitDisplay(B | C | D | F)
	DisplayFive  = DigitDisplay(A | B | D | F | G)
	DisplaySix   = DigitDisplay(A | C | F)
	DisplaySeven = DigitDisplay(A | B | D | E | F | G)
	DisplayEight = DigitDisplay(A | B | C | D | E | F | G)
	DisplayNine  = DigitDisplay(A | B | C | D | F | G)
)

var Displays = [10]DigitDisplay{
	DisplayZero,
	DisplayOne,
	DisplayTwo,
	DisplayThree,
	DisplayFour,
	DisplayFive,
	DisplaySix,
	DisplaySeven,
	DisplayEight,
	DisplayNine,
}
