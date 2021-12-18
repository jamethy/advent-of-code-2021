package advent15

import (
	"fmt"
	"math"
	"strings"

	"advent2021/util"
	"advent2021/util/mathutil"
)

func Solution(inputFile string) (part1, part2 interface{}) {
	riskLevels := parseRiskLevels(inputFile)

	cache := make(map[int]map[int]int, len(riskLevels))
	for x := range riskLevels {
		cache[x] = make(map[int]int, len(riskLevels))
	}
	score := getMinScoreFrom(riskLevels, 0, 0, -1, -1, cache)

	printCache(cache, riskLevels)

	return score - riskLevels[0][0], 0
}


func getMinScoreFrom(riskLevels [][]int, x, y, fromX, fromY int, cache map[int]map[int]int) int {
	if cached, ok := cache[x][y]; ok {
		return cached
	}
	score := riskLevels[y][x]
	if x == y && x == len(riskLevels)-1 {
		cache[x][y] = score
		return score
	}

	right, left, down, up := math.MaxInt, math.MaxInt, math.MaxInt, math.MaxInt
	if x+1 != fromX && x+1 < len(riskLevels) {
		right = getMinScoreFrom(riskLevels, x+1, y, x, y, cache)
	}
	if x != 0 && x-1 != fromX && x-1 < len(riskLevels) {
		left = getMinScoreFrom(riskLevels, x-1, y, x, y, cache)
	}
	if y+1 != fromY && y+1 < len(riskLevels) {
		down = getMinScoreFrom(riskLevels, x, y+1, x, y, cache)
	}
	if y != 0 && y-1 != fromY && y-1 < len(riskLevels) {
		up = getMinScoreFrom(riskLevels, x, y-1, x, y, cache)
	}
	score += mathutil.MinInt(right, left, down, up)
	cache[x][y] = score
	return score
}

func printCache(cache map[int]map[int]int, riskLevels [][]int) {

	cacheValues := make([][]int, len(riskLevels))
	for x := range riskLevels {
		cacheValues[x] = make([]int, len(riskLevels))
	}
	for x, col := range cache {
		for y, v := range col {
			cacheValues[x][y] = v
		}
	}

	for y, r := range cacheValues {
		for x := range r {
			s := fmt.Sprintf("%d ", cacheValues[x][y])
			if len(s) == 2 {
				s = "0" + s
			}
			print(s)
		}
		println()
	}
}

//func Solution(inputFile string) (part1, part2 interface{}) {
//	riskLevels := parseRiskLevels(inputFile)
//	dirs := make([]rune, 0, (len(riskLevels) - 1) * 2)
//	for i := 0; i < len(riskLevels) - 1; i++ {
//		dirs = append(dirs, 'r', 'd')
//	}
//
//	minScore := math.MaxInt
//	Permute(dirs, func(permutation []rune) {
//		score := 0
//		x, y := 0, 0
//		for i, dir := range permutation {
//			if dir == 'r' {
//				x++
//			} else {
//				y++
//			}
//			score += riskLevels[x][y]
//			if score >= minScore {
//				break
//			}
//		}
//		if score < minScore {
//			minScore = score
//		}
//	})
//
//	return minScore, 0
//}
//
//func Permute(dirs []rune, fn func([]rune)) {
//	heapsAlgorithm(dirs, len(dirs), fn)
//}
//
//func heapsAlgorithm(dirs []rune, k int, fn func([]rune)) {
//	if k == 1 {
//		fn(dirs)
//	} else {
//		heapsAlgorithm(dirs, k - 1, fn)
//		for i := 0; i <= k - 1; i++ {
//			if k % 2 == 0 {
//				dirs[i], dirs[k-1] = dirs[k-1], dirs[i]
//			} else {
//				dirs[0], dirs[k-1] = dirs[k-1], dirs[0]
//			}
//			heapsAlgorithm(dirs, k - 1, fn)
//		}
//	}
//}

//type PermutationGenerator struct {
//	Collection []rune
//	Waiting    chan bool
//	Stop       chan bool
//}
//
//func (p PermutationGenerator) Start(col []rune) {
//	p.Collection = col
//	go func() {
//		p.permute(p.Collection, 0, len(col)-1)
//	}()
//}
//
//func (p PermutationGenerator) permute(dirs []rune, l, r int) {
//	if l == r {
//		// permutation
//	} else {
//		for i := l; i <= r; i++ {
//			dirs[l], dirs[r] = dirs[r], dirs[l]
//			p.permute(dirs, l+1, r)
//			dirs[l], dirs[r] = dirs[r], dirs[l]
//		}
//	}
//}

func parseRiskLevels(inputFile string) [][]int {
	lines := util.ReadFile(inputFile)
	res := make([][]int, 0, len(lines))
	for _, line := range lines {
		if line == "" {
			continue
		}
		lineInts := util.StringsToInts(strings.Split(line, ""))
		res = append(res, lineInts)
	}

	return res
}
