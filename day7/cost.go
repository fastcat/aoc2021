package day7

func costTo(
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
) (target, cost int) {
	min, max := posRange(crabs)
	bestTarget, bestCost := min, costTo(crabs, min)
	for i := min + 1; i <= max; i++ {
		if thisCost := costTo(crabs, i); thisCost < bestCost {
			bestTarget, bestCost = i, thisCost
		}
	}
	return bestTarget, bestCost
}
