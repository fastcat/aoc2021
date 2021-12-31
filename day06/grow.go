package day06

import "fmt"

const recycleDelay = 6
const maxDelay = recycleDelay + 2

// number of fish at each growth delay
type GrowthState = [maxDelay + 1]int64

func InitState(input []int) GrowthState {
	var state GrowthState
	for i, val := range input {
		if val > maxDelay {
			panic(fmt.Errorf("bad input %d: %d > %d", i, val, maxDelay))
		}
		state[val]++
	}
	return state
}

func Grow(state GrowthState) GrowthState {
	// births = number at 0 delay, everything else moves down
	births := state[0]
	copy(state[:], state[1:])
	// those that gave birth recycle at an intermediate delay
	state[recycleDelay] += births
	// the new children are out at the max delay
	state[maxDelay] = births
	return state
}

func Total(state GrowthState) int64 {
	total := int64(0)
	for _, n := range state {
		total += n
	}
	return total
}
