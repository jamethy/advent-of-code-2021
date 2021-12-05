package advent01

import (
	"advent2021/util"
)

func Solution(inputFile string) (part1, part2 interface{}) {

	strLines := util.ReadFile(inputFile)
	lines := util.StringsToInts(strLines)

	var part1Count int
	var part2Count int
	for i := range lines {
		if i == 0 {
			continue
		}
		if lines[i] > lines[i-1] {
			part1Count++
		}

		if i >= 3 {
			prev := lines[i-3] + lines[i-2] + lines[i-1]
			this := lines[i-2] + lines[i-1] + lines[i-0]
			if this > prev {
				part2Count++
			}
		}
	}
	return part1Count, part2Count
}
