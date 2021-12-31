package day07

func linearCost(
	crabs []int,
	target int,
) int {
	cost := 0
	for _, pos := range crabs {
		if pos < target {
			cost += target - pos
		} else {
			cost += pos - target
		}
	}
	return cost
}

func quadraticCost(
	crabs []int,
	target int,
) int {
	cost := 0
	for _, pos := range crabs {
		var delta int
		if pos < target {
			delta = target - pos
		} else {
			delta = pos - target
		}
		cost += delta * (delta + 1) / 2
	}
	return cost
}

func posRange(
	crabs []int,
) (min, max int) {
	min, max = crabs[0], crabs[0]
	for _, pos := range crabs[1:] {
		if pos < min {
			min = pos
		}
		if pos > max {
			max = pos
		}
	}
	return
}

func cheapestNaive(
	crabs []int,
	coster func([]int, int) int,
) (target, cost int) {
	min, max := posRange(crabs)
	bestTarget, bestCost := min, coster(crabs, min)
	for i := min + 1; i <= max; i++ {
		if thisCost := coster(crabs, i); thisCost < bestCost {
			bestTarget, bestCost = i, thisCost
		}
	}
	return bestTarget, bestCost
}

// Could make a more cpu-efficient evaluator by remembering the cost for each
// crab at the evaluated position, and the incrementally adjusting that when we
// move by one -- that would be 2 adds per crab, instead of a bunch of
// multiplies and adds. But the naive mode is fast enough for the scale of
// sample data to still only take a couple milliseconds.

// An even better solution for big datasets might be to turn each crab into a
// polynomial (all quadratic) representing is cost as a function of target, and
// then find the minimum of it. However, while it seems that the cost should be
// a quadratic polynomial, the usage of absolute value in the delta means it
// isn't. The positive side (for initial p=0) ix x*(x+1)/2, but the negative
// side is x*(x-1)/2, and thus there is a cusp at x=0 (or x=p) and the
// polynomial expansion is not finite, and this strategy doesn't work in any
// simple form.
