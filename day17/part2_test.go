package day17

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart2Example(t *testing.T) {
	target := Rect{20, 30, -10, -5}
	n := Probe{}.NumVelocitiesThatHit(target)
	assert.Equal(t, 112, n)
}

func TestPart2Challege(t *testing.T) {
	target := Rect{265, 287, -103, -58}
	n := Probe{}.NumVelocitiesThatHit(target)
	t.Logf("challenge options = %d", n)
}
