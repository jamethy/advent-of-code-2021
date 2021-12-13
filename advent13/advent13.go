package advent13

import (
	"fmt"
	"strconv"
	"strings"

	"advent2021/util"
)

func Solution(inputFile string) (part1, part2 interface{}) {
	parts := util.ReadFileSplitBy(inputFile, "\n\n")
	dots := parseDots(strings.Split(parts[0], "\n"))
	folds := parseFolds(strings.Split(parts[1], "\n"))

	var afterFirstFold int

	//for i, fold := range folds {
	dots = foldDots(dots, folds[0])
	dots.Print()

	//if i == 0 {
	afterFirstFold = dots.Count()
	//}
	//}

	return afterFirstFold, 0
}

func foldDots(input Dots, fold Fold) Dots {
	maxX, maxY := input.Max()
	var out Dots

	if fold.Direction == "y" {
		// directly copy 0, 0 -> maxX, maxY/2
		out = input.DuplicateTo(maxX, fold.Number)
		// apply mirrored 0, maxY/2+1 -> maxX, maxY
		for x := 0; x <= maxX; x++ {
			for y := fold.Number + 1; y <= maxY; y++ {
				distFromFold := y - fold.Number
				mirroredY := fold.Number - distFromFold

				if input.Get(x, y) {
					out.Set(x, mirroredY, true)
				}
			}
		}
	} else {
		// directly copy 0, 0 -> maxX/2, maxY
		out = input.DuplicateTo(fold.Number, maxY)
		// apply mirrored maxX/2+1, 0 -> maxX, maxY
		for x := fold.Number + 1; x <= maxX; x++ {
			for y := 0; y <= maxY; y++ {
				distFromFold := x - fold.Number
				mirroredX := fold.Number - distFromFold

				if input.Get(x, y) {
					out.Set(mirroredX, y, true)
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

func (d Dots) DuplicateTo(x, y int) Dots {
	res := Dots{}
	for i := 0; i <= x; i++ {
		for j := 0; j <= y; j++ {
			res.Set(i, j, d.Get(i, j))
		}
	}
	return res
}

func (d Dots) Set(x, y int, value bool) {
	if _, ok := d[x]; !ok {
		d[x] = make(map[int]bool)
	}
	d[x][y] = value
}

func (d Dots) Get(x, y int) bool {
	if _, ok := d[x]; !ok {
		return false
	}
	return d[x][y]
}

func (d Dots) Max() (x, y int) {
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

func (d Dots) Print() {

	sb := strings.Builder{}

	maxX, maxY := d.Max()
	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			if d.Get(x, y) {
				sb.WriteRune('#')
			} else {
				sb.WriteRune('.')
			}
		}
		sb.WriteRune('\n')
	}

	fmt.Println(sb.String())
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

		dots.Set(x, y, true)
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
