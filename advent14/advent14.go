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

	part1 = iterateAndGetScore(template, rules, 10)
	//part2 = iterateAndGetScore(template, rules, 40)

	return part1, part2
}

func iterateAndGetScore(template string, rules map[string]uint8, iterations int) int {
	totalCounts := make(map[string]int)
	for _, char := range template {
		totalCounts[string(char)] += 1
	}
	for i := 1; i < len(template); i++ {
		addCounts(uint8sToString(template[i-1], template[i]), rules, iterations, totalCounts)
	}

	min, max := math.MaxInt, 0
	for _, v := range totalCounts {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return max - min
}

func addCounts(str string, rules map[string]uint8, remaining int, counts map[string]int) {
	if remaining == 0 {
		return
	}
	remaining--

	insert := string(rules[str])
	counts[insert] += 1
	addCounts(uint8sToString(str[0], insert[0]), rules, remaining, counts)
	addCounts(uint8sToString(insert[0], str[1]), rules, remaining, counts)
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
