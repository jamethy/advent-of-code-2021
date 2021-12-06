package advent06

import "advent2021/util"

func Solution(inputFile string) (part1, part2 interface{}) {
	lines := util.ReadFile(inputFile)
	fish := util.ParseIntList(lines[0], ",")

	for i := 0; i < 80; i++ {
		newFish := make([]int, 0)
		for j, f := range fish {
			if f == 0 {
				fish[j] = 6
				newFish = append(newFish, 8)
			} else {
				fish[j] = f - 1
			}
		}
		fish = append(fish, newFish...)
	}

	return len(fish), 0
}
