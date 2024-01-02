package day13

import (
	"aoc2023/lib"
)

func CheckReflect(input [][]byte, lower int, wantDiffs int) bool {
	var diffs int
	upper := lower - 1
	for {
		if upper < 0 || lower == len(input) {
			break
		}
		diffs += DiffCount(input[upper], input[lower])
		upper--
		lower++
	}

	return wantDiffs == diffs
}

func FindReflection(input [][]byte, wantDiffs int) int {
	for index := 1; index < len(input); index++ {
		if CheckReflect(input, index, wantDiffs) {
			return index
		}
	}

	return -1
}

func PatternValue(input [][]byte, diffs int) int {
	if value := FindReflection(input, diffs); value > 0 {
		return 100 * value
	}

	if value := FindReflection(lib.Transpose(input), diffs); value > 0 {
		return value
	}

	panic("unreachable")
}

func ParseInput(input [][]byte, diffs int) (total int) {
	var buf [][]byte

	for _, row := range input {
		if len(row) == 0 {
			total += PatternValue(buf, diffs)
			buf = nil
			continue
		}

		buf = append(buf, row)
	}

	return total + PatternValue(buf, diffs) // last item
}

func DiffCount[T comparable](a, b []T) (count int) {
	for index, n := range a {
		if n != b[index] {
			count++
		}
	}

	return
}
