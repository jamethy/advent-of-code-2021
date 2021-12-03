package advent03

import (
	"strconv"

	"advent2021/util"
)

func Solution(inputFile string) (part1, part2 interface{}) {
	lines, bitLength := readAsBinaryInts(inputFile)

	var gamma uint
	var i uint
	for i = 0; i < bitLength; i++ {
		columnSum := getColumnSum(lines, i)
		if columnSum > len(lines)/2 {
			gamma = setBit(gamma, i)
		}
	}

	eps := flipAllBits(gamma, bitLength)

	oxygen := getRating(lines, bitLength, true)
	co2 := getRating(lines, bitLength, false)

	return int(gamma) * int(eps), int(oxygen) * int(co2)
}

func getColumnSum(lines []uint, pos uint) int {
	res := 0
	for _, line := range lines {
		if isBitSet(line, pos) {
			res++
		}
	}
	return res
}

func getRating(lines []uint, bitLength uint, one bool) uint {

	var i uint
	for i = uint(bitLength) - 1; i >= 0; i-- {
		if len(lines) == 1 {
			return lines[0]
		} else if len(lines) == 2 {
			firstSet, secondSet := isBitSet(lines[0], i), isBitSet(lines[1], i)
			if firstSet != secondSet {
				if firstSet == one {
					return lines[0]
				} else {
					return lines[1]
				}
			}
		}
		sum := getColumnSum(lines, i)
		bitIsSet := one == (sum > len(lines)/2)
		lines = filter(lines, func(u uint) bool {
			return isBitSet(u, i) == bitIsSet
		})
	}
	return 0
}

func filter(lines []uint, f func(uint) bool) []uint {
	res := make([]uint, 0)
	for _, line := range lines {
		if f(line) {
			res = append(res, line)
		}
	}
	return res
}

func isBitSet(i uint, pos uint) bool {
	return i&(1<<pos) > 0
}

func setBit(i uint, pos uint) uint {
	i |= 1 << pos
	return i
}

func flipAllBits(i uint, bits uint) uint {
	maxVal := uint(1)<<bits - 1
	return i ^ maxVal
}

func readAsBinaryInts(inputFile string) ([]uint, uint) {
	lines := util.ReadFile(inputFile)
	bitSize := uint(len(lines[0]))

	res := make([]uint, 0, len(lines))

	for _, s := range lines {
		v, err := strconv.ParseUint(s, 2, int(bitSize))
		if err != nil {
			continue
		}
		res = append(res, uint(v))
	}
	return res, bitSize
}
