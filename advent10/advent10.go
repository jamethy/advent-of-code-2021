package advent10

import (
	"sort"
	"strings"

	"advent2021/util"
)

func Solution(inputFile string) (part1, part2 interface{}) {
	lines := util.ReadFile(inputFile)

	part1Score := 0
	part2Scores := make([]int, 0)
	for _, line := range lines {
		completion, corruptedRune, isCorrupted := complete(line)
		if isCorrupted {
			part1Score += illegalCharacterScore[corruptedRune]
		} else {
			completionScore := 0
			for _, c := range completion {
				completionScore *= 5
				completionScore += completionCharacterScore[c]
			}
			part2Scores = append(part2Scores, completionScore)
		}
	}

	sort.Ints(part2Scores)

	return part1Score, part2Scores[len(part2Scores)/2]
}

func complete(line string) ([]rune, rune, bool) {
	stack := make([]rune, 0, len(line)/3)

	for _, c := range line {
		isChunkStart := strings.ContainsRune("([{<", c)
		if isChunkStart {
			stack = append(stack, c)
		} else if stack[len(stack)-1] == oppositeRune(c) {
			stack = stack[:len(stack)-1]
		} else {
			// is corrupted
			return nil, c, true
		}
	}

	// reverse stack
	completion := make([]rune, len(stack))
	for i := range stack {
		completion = append(completion, oppositeRune(stack[len(stack)-1-i]))
	}

	return completion, '?', false
}

var illegalCharacterScore = map[rune]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

var completionCharacterScore = map[rune]int{
	')': 1,
	']': 2,
	'}': 3,
	'>': 4,
}

func oppositeRune(r rune) rune {
	switch r {
	case ')':
		return '('
	case ']':
		return '['
	case '}':
		return '{'
	case '>':
		return '<'
	case '(':
		return ')'
	case '[':
		return ']'
	case '{':
		return '}'
	case '<':
		return '>'
	default:
		panic("unaccounted for rune")
	}
}
