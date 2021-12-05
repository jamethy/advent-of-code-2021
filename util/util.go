package util

import (
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func Panic(err error) {
	if err != nil {
		panic(err)
	}
}

func StringsToInts(str []string) []int {
	ret := make([]int, 0, len(str))
	for _, s := range str {
		if i, err := strconv.Atoi(s); err == nil {
			ret = append(ret, i)
		}
	}
	return ret
}

func ParseIntList(str, sep string) []int {
	parts := strings.Split(str, sep)
	return StringsToInts(parts)
}

func ReadFile(name string) []string {
	return ReadFileSplitBy(name, "\n")
}

func ReadFileSplitBy(name, delimiter string) []string {
	f, err := os.Open(name)
	if err != nil {
		panic("can't open file " + err.Error())
	}
	d, err := ioutil.ReadAll(f)
	if err != nil {
		panic("can't read file " + err.Error())
	}

	return strings.Split(string(d), delimiter)
}

func FlipString(str string) string {
	str2 := ""
	for i := len(str) - 1; i >= 0; i-- {
		str2 += string(str[i])
	}
	return str2
}


func EqualIntSlice(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}