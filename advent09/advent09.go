package advent09

import (
	"strconv"
	"strings"

	"advent2021/util"
)

type Grid [][]uint8

func Solution(inputFile string) (part1, part2 interface{}) {
	grid := parseGrid(inputFile)

	riskSum := 0
	for i, row := range grid {
		for j := range row {
			v, lowPoint := grid.isLowPoint(i, j)
			if lowPoint {
				riskSum += int(v) + 1
			}
		}
	}

	return riskSum, 0
}

func (g Grid) get(i, j int) (uint8, bool) {
	if i < 0 || i >= len(g) {
		return 0, false
	}
	if j < 0 || j >= len(g[i]) {
		return 0, false
	}
	return g[i][j], true
}

func (g Grid) isLowPoint(i, j int) (uint8, bool) {
	v, ok := g.get(i, j)
	if !ok {
		return v, false
	}

	lowPointNeighbors := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for _, dir := range lowPointNeighbors {
		n, ok := g.get(i+dir[0], j+dir[1])
		if ok && n <= v {
			return v, false
		}
	}

	return v, true
}

func parseGrid(inputFile string) Grid {
	lines := util.ReadFile(inputFile)
	width := len(lines[0])

	grid := make(Grid, 0, len(lines))

	for _, line := range lines {
		if len(line) < width {
			continue
		}
		row := make([]uint8, width)
		digits := strings.Split(line, "")
		for i, digit := range digits {
			value, err := strconv.ParseUint(digit, 10, 8)
			if err == nil {
				row[i] = uint8(value)
			}
		}

		grid = append(grid, row)
	}

	return grid
}
