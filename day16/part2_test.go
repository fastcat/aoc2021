package day16

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPart2Examples(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{"C200B40A82", 3},
		{"04005AC33890", 54},
		{"880086C3E88112", 7},
		{"CE00C43D881120", 9},
		{"D8005AC2A8F0", 1},
		{"F600BC2D8F", 0},
		{"9C005AC2F8F0", 0},
		{"9C0141080250320F1802104A08", 1},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			b, err := DecodeHex(tt.input)
			require.NoError(t, err)
			p := ParsePacket(b)
			assert.Equal(t, tt.want, p.Value())
		})
	}
}

func TestPart2Challenge(t *testing.T) {
	b, err := DecodeHex(strings.TrimSpace(challengeInput))
	require.NoError(t, err)
	packets := ParsePackets(b)
	assert.Len(t, packets, 1)
	t.Logf("remainder = %s", b.String())
	t.Logf("value = %d", packets[0].Value())
}
