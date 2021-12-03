package advent01

import (
	"strconv"

	"advent2021/util"
)

func Solution(inputFile string) (part1, part2 interface{}) {

	lines := fileAsInts(inputFile)

	var count int
	for i, depth := range lines[1:] {
		if depth > lines[i] {
			count++
		}
	}
	return count, 0
}

func fileAsInts(inputFile string) []int {
	lines := util.ReadFile(inputFile)
	ints := make([]int, 0, len(lines))
	for _, s := range lines {
		i, err := strconv.Atoi(s)
		if err != nil {
			continue
		}
		ints = append(ints, i)
	}
	return ints
}
