package day8

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseValue(t *testing.T) {
	tests := []struct {
		pattern   string
		want      Value
		assertion assert.ErrorAssertionFunc
	}{
		{"a", A, nil},
	}
	for _, tt := range tests {
		t.Run(tt.pattern, func(t *testing.T) {
			got, err := ParseValue(tt.pattern)
			if tt.assertion == nil {
				tt.assertion = assert.NoError
			}
			tt.assertion(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
