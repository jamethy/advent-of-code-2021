package advent05

import (
	"strconv"
	"strings"

	"advent2021/util"
)

func Solution(inputFile string) (part1, part2 interface{}) {
	lines := parseInput(inputFile)
	nonDiagonal := filterNonDiagonal(lines)

	var maxX, maxY int
	for _, line := range lines {
		maxX = util.MaxInt(maxX, line.Start.X, line.End.X)
		maxY = util.MaxInt(maxY, line.Start.Y, line.End.Y)
	}

	var part1Count int
	for x := 0; x <= maxX; x++ {
		for y := 0; y <= maxY; y++ {
			var intersectionCount int
			for _, line := range nonDiagonal {
				if line.IntersectsNonDiagonal(x, y) {
					intersectionCount++
					if intersectionCount >= 2 {
						break
					}
				}
			}
			if intersectionCount >= 2 {
				part1Count++
			}
		}
	}

	return part1Count, 0
}

type (
	Point struct {
		X int
		Y int
	}

	Line struct {
		Start Point
		End   Point
	}
)

func (l Line) IntersectsNonDiagonal(x, y int) bool {
	if l.Start.X == l.End.X {
		if x != l.Start.X {
			return false
		}
		return isBetweenInclusive(y, l.Start.Y, l.End.Y)
	} else { // is vertical
		if y != l.Start.Y {
			return false
		}
		return isBetweenInclusive(x, l.Start.X, l.End.X)
	}
}

func isBetweenInclusive(v, a, b int) bool {
	if a < b {
		return v >= a && v <= b
	} else {
		return v <= a && v >= b
	}
}

func filterNonDiagonal(lines []Line) []Line {
	filtered := make([]Line, 0, len(lines))
	for _, line := range lines {
		if line.Start.X == line.End.X || line.Start.Y == line.End.Y {
			filtered = append(filtered, line)
		}
	}

	return filtered
}

func parseInput(inputFile string) []Line {
	fileLines := util.ReadFile(inputFile)
	lines := make([]Line, 0, len(fileLines))

	for _, fileLine := range fileLines {
		if fileLine != "" {
			lines = append(lines, parseLine(fileLine))
		}
	}
	return lines
}

// assumes valid
func parseLine(fileLine string) Line {
	parts := strings.Split(fileLine, " -> ")
	if len(parts) != 2 {
		panic("not two parts")
	}
	return Line{
		Start: parsePoint(parts[0]),
		End:   parsePoint(parts[1]),
	}
}

// assumes valid
func parsePoint(str string) Point {
	parts := strings.Split(str, ",")
	if len(parts) != 2 {
		panic("not two parts")
	}
	x, err := strconv.Atoi(parts[0])
	util.Panic(err)
	y, err := strconv.Atoi(parts[1])
	util.Panic(err)

	return Point{X: x, Y: y}
}
