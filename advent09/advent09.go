package advent09

import (
	"sort"
	"strconv"
	"strings"

	"advent2021/util"
)

type Grid [][]uint8

func Solution(inputFile string) (part1, part2 interface{}) {
	grid := parseGrid(inputFile)

	riskSum := 0
	basinSizes := make([]int, 0)
	for i, row := range grid {
		for j := range row {
			v, lowPoint := grid.isLowPoint(i, j)
			if lowPoint {

				// part 1
				riskSum += int(v) + 1

				// part 2
				basinSizes = append(basinSizes, grid.getBasinSize(i, j))
			}
		}
	}

	sort.Ints(basinSizes)
	basinSizes = basinSizes[len(basinSizes)-3:]

	return riskSum, basinSizes[0] * basinSizes[1] * basinSizes[2]
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

var neighborDirs = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func (g Grid) isLowPoint(i, j int) (uint8, bool) {
	v, ok := g.get(i, j)
	if !ok {
		return v, false
	}

	for _, dir := range neighborDirs {
		n, ok := g.get(i+dir[0], j+dir[1])
		if ok && n <= v {
			return v, false
		}
	}

	return v, true
}

type CoordinateSet [][]int

func (c *CoordinateSet) Add(i, j int) {
	if !c.Has(i, j) {
		*c = append(*c, []int{i, j})
	}
}

func (c *CoordinateSet) Has(i, j int) bool {
	for _, existing := range *c {
		if existing[0] == i && existing[1] == j {
			return true
		}
	}
	return false
}

func (g Grid) getBasinSize(i, j int) int {
	coordinateSet := make(CoordinateSet, 0)
	g.addBasinNeighbors(i, j, &coordinateSet)
	return len(coordinateSet)
}

func (g Grid) addBasinNeighbors(i, j int, c *CoordinateSet) {
	v, ok := g.get(i, j)
	if !ok || v == 9 {
		return
	}
	c.Add(i, j)

	for _, dir := range neighborDirs {
		ni, nj := i+dir[0], j+dir[1]
		if c.Has(ni, nj) {
			continue
		}

		n, ok := g.get(ni, nj)
		if ok && n != 9 {
			g.addBasinNeighbors(ni, nj, c)
		}
	}
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
