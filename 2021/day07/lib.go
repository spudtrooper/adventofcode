package day07

import (
	"math"

	"github.com/spudtrooper/adventofcode/common/ints"
	"github.com/spudtrooper/adventofcode/common/must"
	"github.com/thomaso-mirodin/intmath/intgr"
)

func Part1(input string) int {
	crabs := must.SplitInts(must.ReadAllFile(input), ",")

	cost := func(pos int) int {
		var cost int
		for _, c := range crabs {
			cost += intgr.Abs(pos - c)
		}
		return cost
	}

	min, max := ints.ArrayMin(crabs), ints.ArrayMax(crabs)
	minCost := math.MaxInt
	for pos := min; pos <= max; pos++ {
		minCost = intgr.Min(minCost, cost(pos))
	}

	return minCost

}

func Part1Func(input string) int {
	arr := must.SplitInts(must.ReadAllFile(input), ",")
	crabs := ints.ArrayVec(&arr)

	cost := func(pos int) int {
		return crabs.Iter().Map(func(x int) int { return intgr.Abs(pos - x) }).Sum()
	}

	minCost := ints.FromRange(crabs.Iter().Min(), crabs.Iter().Max()).Reduce(math.MaxInt, func(base, x int) int {
		return intgr.Min(base, cost(x))
	})

	return minCost

}

func Part2(input string) int {
	crabs := must.SplitInts(must.ReadAllFile(input), ",")

	cost := func(pos int) int {
		var cost int
		for _, c := range crabs {
			var total int
			for i, diff := 1, intgr.Abs(pos-c); i <= diff; i++ {
				total += i
			}
			cost += total
		}
		return cost
	}

	min, max := ints.ArrayMin(crabs), ints.ArrayMax(crabs)
	minCost := math.MaxInt
	for pos := min; pos <= max; pos++ {
		minCost = intgr.Min(minCost, cost(pos))
	}

	return minCost

}

func Part2Func(input string) int {
	arr := must.SplitInts(must.ReadAllFile(input), ",")
	crabs := ints.ArrayVec(&arr)

	cost := func(pos int) int {
		return crabs.Iter().Map(func(x int) int {
			return ints.FromRange(1, intgr.Abs(pos-x)).Sum()
		}).Sum()
	}

	minCost := ints.FromRange(crabs.Iter().Min(), crabs.Iter().Max()).Reduce(math.MaxInt, func(base, x int) int {
		return intgr.Min(base, cost(x))
	})

	return minCost

}
