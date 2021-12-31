package day07

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuadraticCostExample(t *testing.T) {
	tests := []struct {
		name   string
		crabs  []int
		target int
		want   int
	}{
		{"1", []int{0}, 1, 1},
		{"2", []int{0}, 2, 3},
		{"3", []int{0}, 3, 6},
		{"example@5", exampleInput, 5, 168},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, quadraticCost(tt.crabs, tt.target))
		})
	}
}

func TestCheapestNaiveQuadraticExample(t *testing.T) {
	target, cost := cheapestNaive(exampleInput, quadraticCost)
	assert.Equal(t, 5, target)
	assert.Equal(t, 168, cost)
}

func TestCheapestNaiveQuadraticChallenge(t *testing.T) {
	target, cost := cheapestNaive(challengeInput, quadraticCost)
	t.Logf("cheapest is %d @ %d", target, cost)
}
