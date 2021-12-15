package advent14

import (
	"math"
	"strings"

	"advent2021/util"
)

func Solution(inputFile string) (part1, part2 interface{}) {
	parts := util.ReadFileSplitBy(inputFile, "\n\n")
	template := parts[0]
	rules := parseInsertionRules(parts[1])

	for i := 0; i < 10; i++ {
		var newTemplate string
		for pos := range template {
			if pos == 0 {
				newTemplate = uint8sToString(template[pos])
				continue
			}
			twoChars := uint8sToString(template[pos-1], template[pos])
			if insert, ok := rules[twoChars]; ok {
				newTemplate += uint8sToString(insert, template[pos])
			} else {
				newTemplate += uint8sToString(template[pos])
			}
		}
		template = newTemplate
	}

	counts := countLetters(template)

	min, max := math.MaxInt, 0
	for _, v := range counts {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}

	return max - min, 0
}

func countLetters(str string) map[uint8]int {
	m := make(map[uint8]int, 0)
	for _, r := range str {
		m[uint8(r)] += 1
	}
	return m
}

func uint8sToString(uint8s ...uint8) string {
	return string(uint8s)
}

func parseInsertionRules(str string) map[string]uint8 {
	rules := make(map[string]uint8, 0)
	lines := strings.Split(str, "\n")
	for _, line := range lines {
		parts := strings.Split(line, " -> ")
		if len(parts) != 2 {
			continue
		}
		rules[parts[0]] = parts[1][0]
	}
	return rules
}
