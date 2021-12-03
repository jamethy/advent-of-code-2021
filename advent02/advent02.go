package advent02

import (
	"errors"
	"strconv"
	"strings"

	"advent2021/util"
)

func Solution(inputFile string) (part1, part2 interface{}) {
	lines := util.ReadFile(inputFile)

	var xPos, yPos, aim, yPos2 int
	for _, line := range lines {
		dir, value, err := lineParts(line)
		if err != nil {
			continue
		}
		switch dir {
		case "forward":
			xPos += value
			yPos2 += aim * value
		case "up":
			yPos -= value
			aim -= value
		case "down":
			yPos += value
			aim += value
		}
	}

	return xPos * yPos, xPos * yPos2
}


func lineParts(line string) (string, int, error) {
	parts := strings.Split(line, " ")
	if len(parts) < 2 {
		return "", 0, errors.New("invalid line")
	}
	num, err := strconv.Atoi(parts[1])
	util.Panic(err)
	return parts[0], num, nil
}