package advent04

import (
	"strconv"
	"strings"

	"advent2021/util"
)

type row []int
type board []row

func Solution(inputFile string) (part1, part2 interface{}) {
	parts := util.ReadFileSplitBy(inputFile, "\n\n")

	calls := splitToInts(parts[0], ",")
	boards := parseBoards(parts[1:])

	winningBoard, winningNumber, called := findWinningBoard(boards, calls)

	return winningBoard.calculateScore(winningNumber, called), 0
}

func (b board) calculateScore(winningNumber int, called util.IntSet) int {
	var score int
	for _, r := range b {
		for _, i := range r {
			if !called.Contains(i) {
				score += i
			}
		}
	}
	return score * winningNumber
}

func findWinningBoard(boards []board, calls []int) (board, int, util.IntSet) {
	called := util.IntSet{}
	for _, call := range calls {
		called.Add(call)
		for _, b := range boards {
			if b.hasBingo(called) {
				return b, call, called
			}
		}
	}
	return nil, 0, nil
}

func (r row) hasBingo(s util.IntSet) bool {
	return s.ContainsSlice(r)
}

func (b board) hasBingo(s util.IntSet) bool {
	// check rows
	for _, r := range b {
		if s.ContainsSlice(r) {
			return true
		}
	}

	// check columns
	for i := range b {
		isBingo := true
		for j := range b {
			if !s.Contains(b[j][i]) {
				isBingo = false
				break
			}
		}
		if isBingo {
			return true
		}
	}

	return false
}

func splitToInts(str, sep string) []int {
	parts := strings.Split(str, sep)
	res := make([]int, 0, len(parts))
	for _, p := range parts {
		i, err := strconv.Atoi(p)
		if err != nil {
			continue
		}
		res = append(res, i)
	}
	return res
}

func parseBoards(parts []string) []board {
	boards := make([]board, 0, len(parts))
	for _, p := range parts {
		lines := strings.Split(p, "\n")
		if len(lines) < 5 {
			continue
		}
		var b board
		for _, line := range lines {
			ints := splitToInts(line, " ")
			if len(ints) > 0 {
				b = append(b, ints)
			}
		}
		if len(b) > 0 {
			boards = append(boards, b)
		}
	}
	return boards
}
