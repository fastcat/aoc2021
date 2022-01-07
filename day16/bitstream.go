package day16

import (
	"fmt"
	"strings"
)

type BitStream struct {
	bits []bool
}

func DecodeHex(input string) (*BitStream, error) {
	bits := make([]bool, 0, len(input)*4)
	var v byte
	for i := 0; i < len(input); i++ {
		if _, err := fmt.Sscanf(input[i:i+1], "%1x", &v); err != nil {
			return nil, err
		}
		for j := 0; j < 4; j++ {
			bits = append(bits, v&0b1000 != 0)
			v <<= 1
		}
	}
	return &BitStream{bits}, nil
}

func DecodeBinary(input string) (*BitStream, error) {
	bits := make([]bool, 0, len(input))
	var v byte
	for i := 0; i < len(input); i++ {
		if _, err := fmt.Sscanf(input[i:i+1], "%1b", &v); err != nil {
			return nil, err
		}
		bits = append(bits, v == 1)
	}
	return &BitStream{bits}, nil
}

func (b *BitStream) Len() int { return len(b.bits) }

func (b *BitStream) NextBool() bool {
	ret := b.bits[0]
	b.bits = b.bits[1:]
	return ret
}

func (b *BitStream) NextByte(len int) byte {
	var ret byte
	for i := 0; i < len; i++ {
		ret <<= 1
		if b.bits[i] {
			ret |= 1
		}
	}
	b.bits = b.bits[len:]
	return ret
}
func (b *BitStream) NextInt(len int) int {
	var ret int
	for i := 0; i < len; i++ {
		ret <<= 1
		if b.bits[i] {
			ret |= 1
		}
	}
	b.bits = b.bits[len:]
	return ret
}

func (b *BitStream) Substream(len int) *BitStream {
	ret := &BitStream{b.bits[:len]}
	b.bits = b.bits[len:]
	return ret
}

func (b *BitStream) String() string {
	s := strings.Builder{}
	s.Grow(b.Len())
	for _, v := range b.bits {
		if v {
			s.WriteRune('1')
		} else {
			s.WriteRune('0')
		}
	}
	return s.String()
}
