package advent06

import (
	"strconv"
	"strings"

	"advent2021/util"
)

func Solution(inputFile string) (part1, part2 interface{}) {
	lines := util.ReadFile(inputFile)
	fish := StringsToInts(strings.Split(lines[0], ","))

	part1Total := 0
	for _, f := range fish {
		part1Total += cacheCalcChildren(80+(6-f)) + 1
	}

	part2Total := 0
	for _, f := range fish {
		part2Total += cacheCalcChildren(256+(6-f)) + 1
	}

	return part1Total, part2Total
}

func cacheCalcChildren(days int) int {
	if r, ok := cache[days]; ok {
		return r
	}
	res := calcChildren(days)
	cache[days] = res
	return res
}

func calcChildren(days int) int {
	if days <= 0 {
		return 0
	}
	total := days / 7
	for days > 9 {
		days -= 7
		total += cacheCalcChildren(days - 2)
	}
	return total
}

var cache map[int]int

func init() {
	cache = make(map[int]int, 256)
}

func StringsToInts(str []string) []int {
	ret := make([]int, 0, len(str))
	for _, s := range str {
		if i, err := strconv.Atoi(s); err == nil {
			ret = append(ret, i)
		}
	}
	return ret
}
