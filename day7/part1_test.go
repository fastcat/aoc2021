package day7

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinearCostExample(t *testing.T) {
	tests := []struct {
		target int
		want   int
	}{
		{2, 37},
		{1, 41},
		{3, 39},
		{10, 71},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("target=%d", tt.target), func(t *testing.T) {
			cost := linearCost(exampleInput, tt.target)
			assert.Equal(t, tt.want, cost)
		})
	}
}

func TestCheapestNaiveLinearExample(t *testing.T) {
	target, cost := cheapestNaive(exampleInput, linearCost)
	assert.Equal(t, 2, target)
	assert.Equal(t, 37, cost)
}

func TestCheapestNaiveLinearChallenge(t *testing.T) {
	target, cost := cheapestNaive(challengeInput, linearCost)
	t.Logf("cheapest is %d @ %d", target, cost)
}
