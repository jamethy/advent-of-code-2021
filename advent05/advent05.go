package advent05

import (
	"strconv"
	"strings"

	"advent2021/util"
	"advent2021/util/mathutil"
)

func Solution(inputFile string) (part1, part2 interface{}) {
	lines := parseInput(inputFile)

	var maxX, maxY int
	for _, line := range lines {
		maxX = mathutil.MaxInt(maxX, line.Start.X, line.End.X)
		maxY = mathutil.MaxInt(maxY, line.Start.Y, line.End.Y)
	}

	var part1Count, part2Count int
	for x := 0; x <= maxX; x++ {
		for y := 0; y <= maxY; y++ {
			var nonDiagonalIntersectionCount, diagonalIntersectionCount int
			for _, line := range lines {
				if line.Intersects(x, y) {
					if line.IsDiagonal() {
						diagonalIntersectionCount++
					} else {
						nonDiagonalIntersectionCount++
					}
				}

				// quit early for better performance
				if nonDiagonalIntersectionCount >= 2 {
					break
				}
			}
			if nonDiagonalIntersectionCount >= 2 {
				part1Count++
				part2Count++
			} else if nonDiagonalIntersectionCount+diagonalIntersectionCount >= 2 {
				part2Count++
			}
		}
	}

	return part1Count, part2Count
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

func (p Point) Sub(o Point) Point {
	return Point{
		X: p.X - o.X,
		Y: p.Y - o.Y,
	}
}

func (l Line) Intersects(x, y int) bool {
	lD := l.End.Sub(l.Start)
	d := Point{X: x, Y: y}.Sub(l.Start)

	// check if co-linear
	cross := lD.X*d.Y - d.X*lD.Y
	if cross != 0 {
		return false
	}
	// now make sure it's in the same direction
	dot := lD.X*d.X + lD.Y*d.Y
	if dot < 0 {
		return false
	}

	// now make sure it's not too long
	lDSquared := lD.X*lD.X + lD.Y*lD.Y
	dSquared := d.X*d.X + d.Y*d.Y
	return lDSquared >= dSquared
}

func (l Line) IsDiagonal() bool {
	return l.Start.X != l.End.X && l.Start.Y != l.End.Y
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
