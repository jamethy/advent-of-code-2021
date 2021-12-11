package advent10

import (
	"strings"

	"advent2021/util"
)

func Solution(inputFile string) (part1, part2 interface{}) {
	lines := util.ReadFile(inputFile)

	part1Score := 0
	for _, line := range lines {
		if c, corrupted := isCorrupted(line); corrupted {
			part1Score += illegalCharacterScore[c]
		}
	}

	return part1Score, 0
}

var illegalCharacterScore = map[rune]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
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
		return '?'
	}
}

func isCorrupted(line string) (rune, bool) {
	stack := make([]rune, 0, len(line)/3)

	for _, c := range line {
		isChunkStart := strings.ContainsRune("([{<", c)
		if isChunkStart {
			stack = append(stack, c)
		} else if stack[len(stack)-1] == oppositeRune(c) {
			stack = stack[:len(stack)-1]
		} else {
			return c, true
		}
	}
	return '?', false
}
