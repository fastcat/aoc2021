package day8

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValue_String(t *testing.T) {
	tests := []struct {
		v    Value
		want string
	}{
		{A, "a"},
		{B, "b"},
		{C, "c"},
		{D, "d"},
		{E, "e"},
		{F, "f"},
		{G, "g"},
		{A | C | F, "acf"},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%07b", tt.v), func(t *testing.T) {
			assert.Equal(t, tt.want, tt.v.String())
		})
	}
}

func TestDigitOption_String(t *testing.T) {
	tests := []struct {
		d    DigitOption
		want string
	}{
		{Zero, "0"},
		{One, "1"},
		{Two, "2"},
		{Three, "3"},
		{Four, "4"},
		{Five, "5"},
		{Six, "6"},
		{Seven, "7"},
		{Eight, "8"},
		{Nine, "9"},
		{Zero | Three | Seven, "037"},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%010b", tt.d), func(t *testing.T) {
			assert.Equal(t, tt.want, tt.d.String())
		})
	}
}
