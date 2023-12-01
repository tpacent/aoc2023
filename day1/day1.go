package day1

import (
	"strconv"
	"strings"
)

const (
	Digit0 = '0'
	Digit9 = '9'
)

func digitFunc(r rune) bool {
	return r >= Digit0 && r <= Digit9
}

func MustRecoverValue(s string) int {
	first := strings.IndexFunc(s, digitFunc)
	last := strings.LastIndexFunc(s, digitFunc)
	snumber := string(s[first]) + string(s[last])
	if n, err := strconv.Atoi(string(snumber)); err == nil {
		return n
	}

	panic("unreachable")
}

var figures = map[string]int{
	"0":     0,
	"1":     1,
	"2":     2,
	"3":     3,
	"4":     4,
	"5":     5,
	"6":     6,
	"7":     7,
	"8":     8,
	"9":     9,
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func extractFirstDigit(s string) int {
	return extractDigit(s, strings.Index, func(prev, curr int) bool {
		return curr < prev || prev < 0
	})
}

func extractLastDigit(s string) int {
	return extractDigit(s, strings.LastIndex, func(prev, curr int) bool {
		return curr > prev
	})
}

func extractDigit(s string, indexFunc func(s, substr string) int, okFunc func(prev, curr int) bool) int {
	var (
		index     int
		lastIndex = -1
		lastKey   string
	)

	for key := range figures {
		index = indexFunc(s, key)

		if index < 0 {
			continue
		}

		if okFunc(lastIndex, index) {
			lastIndex = index
			lastKey = key
		}
	}

	if lastIndex < 0 {
		panic("digit not found")
	}

	return figures[lastKey]
}

func MustRecoverSpelledValue(s string) int {
	return extractFirstDigit(s)*10 + extractLastDigit(s)
}
