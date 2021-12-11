package advent08

import (
	"regexp"
	"strconv"
	"strings"

	"advent2021/util"
)

type Line struct {
	SignalPatterns []string // 10
	OutputDigits   []string // 4
}

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

//   aaaa
//  b    c
//  b    c
//   dddd
//  e    f
//  e    f
//   gggg

// from SignalPatterns, determine mapping from jumbled to  proper output
func (l Line) decodeMap() map[string]string {
	remaining := make(map[string]string)
	for _, a := range "abcdefg" {
		remaining[string(a)] = "abcdefg"
	}

	length := 0

	for length != lengthOfMap(remaining) {
		length = lengthOfMap(remaining)

		allFiveCounts := make([]string, 0)
		allSixCounts := make([]string, 0)

		for _, u := range l.SignalPatterns {
			_, rs := getPossibles(u, remaining)
			if len(rs) == 1 {
				filterMatches(remaining, u, rs[0])
			}

			if len(u) == 5 {
				allFiveCounts = append(allFiveCounts, u)
			} else if len(u) == 6 {
				allSixCounts = append(allSixCounts, u)
			}
		}

		filterCommons(allFiveCounts, []string{"acdeg", "acdfg", "abdfg"}, remaining)
		filterCommons(allSixCounts, []string{"abcefg", "abdefg", "abcdfg"}, remaining)
	}

	return remaining
}

func getPossibles(u string, m map[string]string) ([]int, []string) {
	switch len(u) {
	case 2:
		return []int{1}, []string{"cf"}
	case 4:
		return []int{4}, []string{"bcdf"}
	case 5:
		return filterPossibles(u, []int{2, 3, 5}, []string{"acdeg", "acdfg", "abdfg"}, m)
	case 6:
		return filterPossibles(u, []int{0, 6, 9}, []string{"abcefg", "abdefg", "abcdfg"}, m)
	case 3:
		return []int{7}, []string{"acf"}
	case 7:
		return []int{8}, []string{"abcdefg"}
	default:
		panic("no matches possible " + u)
	}
}

func lengthOfMap(m map[string]string) int {
	total := 0
	for _, v := range m {
		total += len(v)
	}
	return total
}

func filterCommons(us []string, strs []string, m map[string]string) {
	commonCoded := "abcdefg"
	for _, u := range us {
		commonCoded = removeNonMatches(commonCoded, u)
	}

	commonUncoded := "abcdefg"
	for _, s := range strs {
		commonUncoded = removeNonMatches(commonUncoded, s)
	}

	filterMatches(m, commonCoded, commonUncoded)
}

func filterPossibles(u string, nums []int, strs []string, m map[string]string) ([]int, []string) {
	ri := make([]int, 0)
	rs := make([]string, 0)

	// check if it has all required
	decodedWithMaybes := ""
	for _, a := range u {
		decodedWithMaybes += m[string(a)]
	}

	for i, uncodedNum := range strs {

		hasAll := true
		for _, a := range uncodedNum {
			if !strings.ContainsRune(decodedWithMaybes, a) {
				hasAll = false
				break
			}
		}

		if hasAll {
			ri = append(ri, nums[i])
			rs = append(rs, strs[i])
		}
	}

	return ri, rs
}

func reverseMap(m map[string]string) map[string]string {
	r := make(map[string]string)
	for k, v := range m {
		for _, a := range v {
			s := r[string(a)]
			if !strings.Contains(s, k) {
				r[string(a)] += k
			}
		}
	}
	return r
}

func determineNumber(u string, m map[string]string) (int, bool) {
	switch len(u) {
	case 5:
		is2, maybe2 := isCode("acdeg", u, m)
		is3, maybe3 := isCode("acdfg", u, m)
		is5, maybe5 := isCode("abdfg", u, m)
		if is2 || (maybe2 && !maybe3 && !maybe5) {
			return 2, true
		}
		if is3 || (!maybe2 && maybe3 && !maybe5) {
			return 3, true
		}
		if is5 || (!maybe2 && !maybe3 && maybe5) {
			return 5, true
		}
		return -1, false

		//2 a c d e g
		//3 a c d f g
		//5 a b d f g
		// all have adg
	case 6:
		is0, maybe0 := isCode("abcefg", u, m)
		is6, maybe6 := isCode("abdefg", u, m)
		is9, maybe9 := isCode("abcdfg", u, m)
		if is0 || (maybe0 && !maybe6 && !maybe9) {
			return 0, true
		}
		if is6 || (!maybe0 && maybe6 && !maybe9) {
			return 6, true
		}
		if is9 || (!maybe0 && !maybe6 && maybe9) {
			return 9, true
		}
		return -1, false
	default:
		return -1, false
	}
}

func isCode(code string, u string, m map[string]string) (yes, maybe bool) {
	yes = true
	for _, a := range code {
		possibles := reversePossibles(m, a)
		if len(possibles) == 1 {
			if !strings.Contains(u, possibles) {
				return false, false
			}
		} else {
			yes = false
			if !strings.ContainsAny(u, possibles) {
				return false, false
			}
		}
	}
	return yes, true
}

func reversePossibles(m map[string]string, a rune) string {
	res := ""
	for k, v := range m {
		if strings.ContainsRune(v, a) {
			res += k
		}
	}
	return res
}

func filterMatches(m map[string]string, input, possibles string) {
	// first, for the letters in input, we limit their remaining by possibles
	for _, a := range input {
		s := string(a)
		m[s] = removeNonMatches(m[s], possibles)
	}

	// next, for the letters not in input, we remove possibles
	for _, a := range removeMatches("abcdefg", input) {
		s := string(a)
		m[s] = removeMatches(m[s], possibles)
	}
}

func removeMatches(str, possibles string) string {
	return regexp.MustCompile("["+possibles+"]").ReplaceAllString(str, "")
}

func removeNonMatches(str, possibles string) string {
	return regexp.MustCompile("[^"+possibles+"]").ReplaceAllString(str, "")
}

func (l Line) decode() int {
	m := l.decodeMap()

	outputStr := ""
	for _, d := range l.OutputDigits {
		i, _ := getPossibles(d, m)
		if len(i) != 1 {
			panic("not right")
		}
		outputStr += strconv.Itoa(i[0])
	}

	n, _ := strconv.Atoi(outputStr)
	return n
}