package advent04

import (
	"strings"

	"advent2021/util"
	"advent2021/util/set"
)

type row []int
type board []row

func Solution(inputFile string) (part1, part2 interface{}) {
	parts := util.ReadFileSplitBy(inputFile, "\n\n")

	calls := util.ParseIntList(parts[0], ",")
	boards := parseBoards(parts[1:])

	winningBoard, winningNumber, called := findWinningBoard(boards, calls)
	losingBoard, losingNumber, losingCalled := findLosingBoard(boards, calls)

	return winningBoard.calculateScore(winningNumber, called), losingBoard.calculateScore(losingNumber, losingCalled)
}

func (b board) calculateScore(winningNumber int, called set.Ints) int {
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

func findWinningBoard(boards []board, calls []int) (board, int, set.Ints) {
	called := set.Ints{}
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

func findLosingBoard(boards []board, calls []int) (board, int, set.Ints) {
	called := set.Ints{}
	for _, call := range calls {
		called.Add(call)

		losingBoards := make([]board, 0)
		for _, b := range boards {
			if !b.hasBingo(called) {
				losingBoards = append(losingBoards, b)
			}
		}

		if len(losingBoards) == 1 {
			return findWinningBoard(losingBoards, calls)
		}
	}
	return nil, 0, nil
}

func (r row) hasBingo(s set.Ints) bool {
	return s.ContainsSlice(r)
}

func (b board) hasBingo(s set.Ints) bool {
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

func parseBoards(parts []string) []board {
	boards := make([]board, 0, len(parts))
	for _, p := range parts {
		lines := strings.Split(p, "\n")
		if len(lines) < 5 {
			continue
		}
		var b board
		for _, line := range lines {
			ints := util.ParseIntList(line, " ")
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
