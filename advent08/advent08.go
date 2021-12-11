package advent08

import (
	"regexp"
	"strconv"
	"strings"

	"advent2021/util"
)

type Line struct {
	SignalPatterns []string // length 10
	OutputDigits   []string // length 4
}

const allLetters = "abcdefg"

func Solution(inputFile string) (part1, part2 interface{}) {
	lines := parseLines(inputFile)

	uniqueOutputCount := 0
	for _, line := range lines {
		for _, digit := range line.OutputDigits {
			if _, ok := uniqueNum(digit); ok {
				uniqueOutputCount++
			}
		}
	}

	outputSum := 0
	for _, line := range lines {
		outputSum += line.decode()
	}

	return uniqueOutputCount, outputSum
}

func parseLines(inputFile string) []Line {
	strs := util.ReadFile(inputFile)
	lines := make([]Line, 0, len(strs))
	for _, str := range strs {
		if !strings.Contains(str, "|") {
			continue
		}
		lines = append(lines, parseLine(str))
	}

	return lines
}

func parseLine(str string) Line {
	parts := strings.Split(str, "|")
	return Line{
		SignalPatterns: strings.Split(strings.TrimSpace(parts[0]), " "),
		OutputDigits:   strings.Split(strings.TrimSpace(parts[1]), " "),
	}
}

// figure out number from unique number of lines
func uniqueNum(str string) (int, bool) {
	switch len(str) {
	case 2:
		return 1, true
	case 4:
		return 4, true
	case 3:
		return 7, true
	case 7:
		return 8, true
	default:
		return -1, false
	}
}

func (l Line) decode() int {
	m := l.decodeMap()

	outputStr := ""
	for _, d := range l.OutputDigits {
		i, _ := getPossibleDecodings(d, m)
		if len(i) != 1 {
			panic("Not the right number of possibles to decode")
		}
		outputStr += strconv.Itoa(i[0])
	}

	n, _ := strconv.Atoi(outputStr)
	return n
}

// from SignalPatterns, determine mapping from jumbled to  proper output
func (l Line) decodeMap() map[rune]string {
	remaining := make(map[rune]string)
	for _, a := range allLetters {
		remaining[a] = allLetters
	}

	for length := -1; length != lengthOfMap(remaining); length = lengthOfMap(remaining) {

		for _, signalPattern := range l.SignalPatterns {
			_, possibleValues := getPossibleDecodings(signalPattern, remaining)
			if len(possibleValues) == 1 {
				filterMatches(remaining, signalPattern, possibleValues[0])
			}

		}

		allFiveCounts := make([]string, 0)
		allSixCounts := make([]string, 0)
		for _, signalPattern := range l.SignalPatterns {
			if len(signalPattern) == 5 {
				allFiveCounts = append(allFiveCounts, signalPattern)
			} else if len(signalPattern) == 6 {
				allSixCounts = append(allSixCounts, signalPattern)
			}
		}

		filterCommons(allFiveCounts, []string{"acdeg", "acdfg", "abdfg"}, remaining)
		filterCommons(allSixCounts, []string{"abcefg", "abdefg", "abcdfg"}, remaining)
	}

	return remaining
}

func getPossibleDecodings(signalPattern string, m map[rune]string) ([]int, []string) {
	switch len(signalPattern) {
	case 2:
		return []int{1}, []string{"cf"}
	case 4:
		return []int{4}, []string{"bcdf"}
	case 5:
		return filterPossibles(signalPattern, []int{2, 3, 5}, []string{"acdeg", "acdfg", "abdfg"}, m)
	case 6:
		return filterPossibles(signalPattern, []int{0, 6, 9}, []string{"abcefg", "abdefg", "abcdfg"}, m)
	case 3:
		return []int{7}, []string{"acf"}
	case 7:
		return []int{8}, []string{allLetters}
	default:
		panic("no matches possible " + signalPattern)
	}
}

func lengthOfMap(m map[rune]string) int {
	total := 0
	for _, v := range m {
		total += len(v)
	}
	return total
}

func filterCommons(signalPattern []string, strs []string, m map[rune]string) {
	commonCoded := allLetters
	for _, signalPattern := range signalPattern {
		commonCoded = removeNonMatches(commonCoded, signalPattern)
	}

	commonUncoded := allLetters
	for _, s := range strs {
		commonUncoded = removeNonMatches(commonUncoded, s)
	}

	filterMatches(m, commonCoded, commonUncoded)
}

func filterMatches(m map[rune]string, coded, decoded string) {
	// first, for the letters in coded, we limit their remaining by decoded
	for _, a := range coded {
		m[a] = removeNonMatches(m[a], decoded)
	}

	// next, for the letters not in coded, we remove decoded
	for _, a := range removeMatches(allLetters, coded) {
		m[a] = removeMatches(m[a], decoded)
	}
}

func removeMatches(str, possibles string) string {
	return regexp.MustCompile("["+possibles+"]").ReplaceAllString(str, "")
}

func removeNonMatches(str, possibles string) string {
	return regexp.MustCompile("[^"+possibles+"]").ReplaceAllString(str, "")
}

func filterPossibles(signalPattern string, nums []int, decodedStrings []string, m map[rune]string) ([]int, []string) {
	filteredDigits := make([]int, 0)
	filteredDecodedStrings := make([]string, 0)

	// check if it has all required
	decodedWithMaybes := ""
	for _, letter := range signalPattern {
		decodedWithMaybes += m[letter]
	}

	for i, decodedString := range decodedStrings {

		hasAll := true
		for _, letter := range decodedString {
			if !strings.ContainsRune(decodedWithMaybes, letter) {
				hasAll = false
				break
			}
		}

		if hasAll {
			filteredDigits = append(filteredDigits, nums[i])
			filteredDecodedStrings = append(filteredDecodedStrings, decodedStrings[i])
		}
	}

	return filteredDigits, filteredDecodedStrings
}

