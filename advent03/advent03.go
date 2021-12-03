package advent03

import (
	"strconv"

	"advent2021/util"
)

func Solution(inputFile string) (part1, part2 interface{}) {
	lines, bitLength := readAsBinaryInts(inputFile)

	columnSums := make([]int, bitLength)
	var i uint

	for _, line := range lines {
		for i = 0; i < bitLength; i++ {
			if isBitSet(line, i) {
				columnSums[i]++
			}
		}
	}

	var gamma uint
	for i = 0; i < bitLength; i++ {
		if columnSums[i] >= len(lines)/2 {
			gamma = setBit(gamma, i)
		}
	}

	eps := flipAllBits(gamma, bitLength)
	return int(gamma) * int(eps), 0
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
