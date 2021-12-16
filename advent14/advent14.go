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

	cache := make(map[string]map[int]map[string]int)
	part1 = iterateAndGetScore(template, rules, 10, cache)
	part2 = iterateAndGetScore(template, rules, 40, cache)

	return part1, part2
}

// Cache hilariously stupid map
type Cache map[string]map[int]map[string]int

func iterateAndGetScore(template string, rules map[string]uint8, iterations int, cache Cache) int {
	totalCounts := make(map[string]int)
	for _, char := range template {
		totalCounts[string(char)] += 1
	}
	for i := 1; i < len(template); i++ {
		cachedCount(uint8sToString(template[i-1], template[i]), rules, iterations, totalCounts, cache)
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

func cachedCount(str string, rules map[string]uint8, remaining int, total map[string]int, cache Cache) {
	if remaining == 0 {
		return
	}
	remaining--

	// check cache
	if cached, ok := cache[str]; ok {
		if values, ok := cached[remaining]; ok {
			for k, v := range values {
				total[k] += v
			}
			return
		}
	} else {
		cache[str] = make(map[int]map[string]int)
	}

	insert := string(rules[str])
	counts := make(map[string]int, 1)
	counts[insert] = 1
	cachedCount(uint8sToString(str[0], insert[0]), rules, remaining, counts, cache)
	cachedCount(uint8sToString(insert[0], str[1]), rules, remaining, counts, cache)

	for k, v := range counts {
		total[k] += v
	}

	cache[str][remaining] = counts
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
