package advent07

import (
	"math"

	"advent2021/util"
	"advent2021/util/mathutil"
)

func Solution(inputFile string) (part1, part2 interface{}) {
	lines := util.ReadFile(inputFile)
	crabs := util.ParseIntList(lines[0], ",")

	minFuelPart1, minFuelPart2 := math.MaxInt, math.MaxInt
	minCrabPos := mathutil.MinInt(crabs...)
	maxCrabPos := mathutil.MaxInt(crabs...)

	for i := minCrabPos; i <= maxCrabPos; i++ {
		fuelPart1, fuelPart2 := 0, 0
		for _, c := range crabs {
			distance := mathutil.AbsInt(i - c)
			fuelPart1 += distance
			fuelPart2 += mathutil.SumOfN(distance)
		}

		if fuelPart1 < minFuelPart1 {
			minFuelPart1 = fuelPart1
		}

		if fuelPart2 < minFuelPart2 {
			minFuelPart2 = fuelPart2
		}
	}
	return minFuelPart1, minFuelPart2
}
