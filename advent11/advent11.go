package advent11

import (
	"strconv"
	"strings"

	"advent2021/util"
)

type Grid [][]uint8

func Solution(inputFile string) (part1, part2 interface{}) {
	grid := parseGrid(inputFile)

	flashOfFirst100 := 0
	stepOfSynchronousFlash := 0

	for i := 0; ; i++ {
		flashCount := grid.stepAndCountFlashes()
		if i < 100 {
			flashOfFirst100 += flashCount
		}

		if flashCount == 100 {
			stepOfSynchronousFlash = i + 1
			break
		}
	}

	return flashOfFirst100, stepOfSynchronousFlash
}

func (g Grid) stepAndCountFlashes() int {
	for i, row := range g {
		for j := range row {
			g.increment(i, j)
		}
	}

	flashCount := 0
	for i, row := range g {
		for j := range row {
			if g[i][j] >= 10 {
				g[i][j] = 0
				flashCount++
			}
		}
	}

	return flashCount
}

func (g Grid) increment(i, j int) {
	if i < 0 || i >= len(g) {
		return
	}
	if j < 0 || j >= len(g[i]) {
		return
	}

	g[i][j] += 1

	// check for flash
	if g[i][j] == 10 {
		for _, dir := range neighborDirs {
			g.increment(i+dir[0], j+dir[1])
		}
	}
}

var neighborDirs = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}, {-1, -1}, {1, 1}, {1, -1}, {-1, 1}}

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
