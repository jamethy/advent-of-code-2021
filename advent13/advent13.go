package advent13

import (
	"strconv"
	"strings"

	"advent2021/util"
)

func Solution(inputFile string) (part1, part2 interface{}) {
	parts := util.ReadFileSplitBy(inputFile, "\n\n")
	dots := parseDots(strings.Split(parts[0], "\n"))
	folds := parseFolds(strings.Split(parts[1], "\n"))

	var afterFirstFold int

	for i, fold := range folds {
		dots = foldDots(dots, fold)

		if i == 0 {
			afterFirstFold = dots.Count()
		}
	}

	return afterFirstFold, "\n" + dots.Print()
}

func foldDots(input Dots, fold Fold) Dots {
	maxX, maxY := input.MaxPoint()
	var out Dots

	if fold.Direction == "y" {
		out = input.DuplicateTo(maxX, fold.Number)
		for x := 0; x <= maxX; x++ {
			for y := fold.Number + 1; y <= maxY; y++ {
				if input.IsMarked(x, y) {
					distFromFold := y - fold.Number
					mirroredY := fold.Number - distFromFold

					out.MarkDot(x, mirroredY)
				}
			}
		}
	} else {
		out = input.DuplicateTo(fold.Number, maxY)
		for x := fold.Number + 1; x <= maxX; x++ {
			for y := 0; y <= maxY; y++ {

				if input.IsMarked(x, y) {
					distFromFold := x - fold.Number
					mirroredX := fold.Number - distFromFold

					out.MarkDot(mirroredX, y)
				}
			}
		}
	}

	return out
}

type Dots map[int]map[int]bool

type Fold struct {
	Direction string
	Number    int
}

func (d Dots) DuplicateTo(maxX, maxY int) Dots {
	res := Dots{}
	for x := 0; x <= maxX; x++ {
		for y := 0; y <= maxY; y++ {
			if d.IsMarked(x, y) {
				res.MarkDot(x, y)
			}
		}
	}
	return res
}

func (d Dots) MarkDot(x, y int) {
	if _, ok := d[x]; !ok {
		d[x] = make(map[int]bool)
	}
	d[x][y] = true
}

func (d Dots) IsMarked(x, y int) bool {
	if _, ok := d[x]; !ok {
		return false
	}
	return d[x][y]
}

func (d Dots) MaxPoint() (x, y int) {
	for j, ys := range d {
		if j > x {
			x = j
		}
		for i := range ys {
			if i > y {
				y = i
			}
		}
	}
	return x, y
}

func (d Dots) Count() int {
	count := 0
	for _, ys := range d {
		for _, v := range ys {
			if v {
				count++
			}
		}
	}
	return count
}

func (d Dots) Print() string {

	sb := strings.Builder{}

	maxX, maxY := d.MaxPoint()
	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			if d.IsMarked(x, y) {
				sb.WriteRune('#')
			} else {
				sb.WriteRune('.')
			}
		}
		sb.WriteRune('\n')
	}

	return sb.String()
}

func parseDots(lines []string) Dots {
	dots := make(Dots)

	for _, line := range lines {
		parts := strings.Split(line, ",")
		if len(parts) != 2 {
			continue
		}
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])

		dots.MarkDot(x, y)
	}
	return dots
}

func parseFolds(lines []string) []Fold {
	folds := make([]Fold, 0)

	for _, line := range lines {
		parts := strings.Split(line, " ")
		if len(parts) != 3 {
			continue
		}
		parts = strings.Split(parts[2], "=")
		i, _ := strconv.Atoi(parts[1])

		folds = append(folds, Fold{
			Direction: parts[0],
			Number:    i,
		})
	}

	return folds
}
