package day13

import (
	"aoc2023/lib"
	"bytes"
)

func CheckReflect(input [][]byte, upper, lower int) bool {
	for {
		if upper < 0 || lower > len(input)-1 {
			break
		}
		if !bytes.Equal(input[upper], input[lower]) {
			return false
		}
		upper--
		lower++
	}

	return true
}

func FindReflection(input [][]byte) int {
	for index := 1; index < len(input); index++ {
		if CheckReflect(input, index-1, index) {
			return index
		}
	}

	return -1
}

func PatternValue(input [][]byte) int {
	if value := FindReflection(input); value > 0 {
		return 100 * value
	}

	if value := FindReflection(lib.Transpose(input)); value > 0 {
		return value
	}

	panic("unreachable")
}

func ParseInput(input [][]byte) (total int) {
	var buf [][]byte

	for _, row := range input {
		if len(row) == 0 {
			total += PatternValue(buf)
			buf = nil
			continue
		}

		buf = append(buf, row)
	}

	total += PatternValue(buf) // last item
	return
}
