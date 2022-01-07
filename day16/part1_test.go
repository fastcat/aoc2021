package day16

import (
	_ "embed"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPart1Examples(t *testing.T) {
	tests := []struct {
		hex       string
		binary    string
		remainder string
		want      Packet
	}{
		{
			"D2FE28",
			"110100101111111000101000",
			"000",
			&Literal{Header{6, 4}, 2021},
		},
		{
			"38006F45291200",
			"00111000000000000110111101000101001010010001001000000000",
			"0000000",
			&Operator{
				Header{1, 6},
				false,
				[]Packet{
					&Literal{Header{6, 4}, 10},
					&Literal{Header{2, 4}, 20},
				},
			},
		},
		{
			"EE00D40C823060",
			"11101110000000001101010000001100100000100011000001100000",
			"00000",
			&Operator{
				Header{7, 3},
				true,
				[]Packet{
					&Literal{Header{2, 4}, 1},
					&Literal{Header{4, 4}, 2},
					&Literal{Header{1, 4}, 3},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.hex, func(t *testing.T) {
			b, err := DecodeHex(tt.hex)
			require.NoError(t, err)
			assert.Equal(t, tt.binary, b.String())
			p := ParsePacket(b)
			assert.Equal(t, tt.remainder, b.String())
			assert.Equal(t, tt.want, p)
		})
	}
}

func TestExampleStructureSums(t *testing.T) {
	tests := []struct {
		hex  string
		want int
	}{
		{"8A004A801A8002F478", 16},
		{"620080001611562C8802118E34", 12},
		{"C0015000016115A2E0802F182340", 23},
		{"A0016C880162017C3686B18A3D4780", 31},
	}
	for _, tt := range tests {
		t.Run(tt.hex, func(t *testing.T) {
			b, err := DecodeHex(tt.hex)
			require.NoError(t, err)
			p := ParsePacket(b)
			assert.Equal(t, tt.want, VersionSum(p))
		})
	}
}

//go:embed challenge.txt
var challengeInput string

func TestPart1Challenge(t *testing.T) {
	b, err := DecodeHex(strings.TrimSpace(challengeInput))
	require.NoError(t, err)
	packets := ParsePackets(b)
	sum := VersionSum(packets...)
	t.Logf("sum = %d", sum)
}
