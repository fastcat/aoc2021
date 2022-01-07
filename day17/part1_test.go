package day17

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1Examples(t *testing.T) {
	target := Rect{20, 30, -10, -5}
	tests := []struct {
		// target Rect
		ivel Vel
		hits bool
		maxY int
	}{
		// for positive initial Y velocity, maxY should always be n(n+1)/2
		// same applies to positivie initial X velocity
		{Vel{7, 2}, true, 3},
		{Vel{6, 3}, true, 6},
		{Vel{9, 0}, true, 0},
		{Vel{17, -4}, false, 0},
		{Vel{6, 9}, true, 45},
		{Vel{6, 10}, false, 55},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			p := Probe{Pos{0, 0}, tt.ivel}
			hits, maxY, _ := p.WillHit(target)
			assert.Equal(t, tt.hits, hits)
			assert.Equal(t, tt.maxY, maxY)
		})
	}

	v1, v2 := Probe{}.OptimalVelocityRange(target)
	h1, maxY, _ := Probe{Vel: v1}.WillHit(target)
	h2, _, _ := Probe{Vel: v2}.WillHit(target)
	assert.True(t, h1)
	assert.True(t, h2)
	assert.Equal(t, 45, maxY)
	assert.Equal(t, Vel{7, 9}, v1)
	assert.Equal(t, Vel{6, 4}, v2)
}

func TestPart1Challenge(t *testing.T) {
	target := Rect{265, 287, -103, -58}
	p := Probe{}
	v1, _ := p.OptimalVelocityRange(target)
	p.Vel = v1
	hits, maxY, at := p.WillHit(target)
	assert.True(t, hits)
	t.Logf("challenge hits at %v after %d from %v", at, maxY, p.Vel)
}
