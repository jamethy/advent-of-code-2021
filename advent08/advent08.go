package advent08

import (
	"strings"

	"advent2021/util"
)

type Line struct {
	SignalPatterns []string // 10
	OutputDigits   []string // 4
}

func Solution(inputFile string) (part1, part2 interface{}) {
	lines := parseLines(inputFile)

	uniqueOutputCount := 0
	for _, line := range lines {
		for _, digit := range line.OutputDigits {
			if _, ok := uniqueNum(digit); ok {
				uniqueOutputCount++
			}
		}
	}
	return uniqueOutputCount, 0
}

func parseLines(inputFile string) []Line {
	strs := util.ReadFile(inputFile)
	lines := make([]Line, 0, len(strs))
	for _, str := range strs {
		parts := strings.Split(str, "|")
		if len(parts) != 2 {
			continue
		}
		lines = append(lines, Line{
			SignalPatterns: strings.Split(parts[0], " "),
			OutputDigits:   strings.Split(parts[1], " "),
		})
	}

	return lines
}

// figure out number from unique number of lines
func uniqueNum(str string) (int, bool) {
	switch len(str) {
	case 2:
		return 1, true
	case 4:
		return 4, true
	case 3:
		return 7, true
	case 7:
		return 8, true
	default:
		return -1, false
	}
}
