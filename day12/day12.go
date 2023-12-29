package day12

import (
	"aoc2023/lib"
	"bytes"
	"errors"
	"strings"
)

const (
	Empty   = '.'
	Damaged = '#'
	Unknown = '?'
)

func FindArrangements(data []byte, spec []int) (indexes [][]int) {
	segmentLen, rest := spec[0], spec[1:]

	for k := 0; k < len(data)-segmentLen+1; k++ {
		if !checkSegment(data, k, segmentLen) {
			continue
		}

		if len(rest) == 0 {
			indexes = append(indexes, []int{k})
			continue
		}

		nextSliceIndex := k + segmentLen + 1
		if nextSliceIndex >= len(data) {
			continue
		}

		for _, sub := range FindArrangements(data[nextSliceIndex:], rest) {
			for idx := range sub {
				sub[idx] += nextSliceIndex
			}
			rec := append([]int{k}, sub...)
			indexes = append(indexes, rec)
		}
	}

	return
}

func checkSegment(data []byte, idx, slen int) bool {
	// check chunk
	if bytes.IndexByte(data[idx:idx+slen], Empty) >= 0 {
		return false
	}
	// check head
	if idx > 0 && data[idx-1] == Damaged {
		return false
	}
	// check tail
	if idx+slen < len(data) && data[idx+slen] == Damaged {
		return false
	}
	return true
}

func FilterValidSolutions(data []byte, spec []int, solutions [][]int) [][]int {
	valids := make([][]int, 0, len(solutions))

	for _, s := range solutions {
		sol := ApplySolution(data, spec, s)
		if bytes.IndexByte(sol, Damaged) < 0 {
			valids = append(valids, s)
		}
	}

	return valids
}

func ApplySolution(data []byte, spec []int, solution []int) []byte {
	out := append([]byte(nil), data...)

	for index, s := range solution {
		segmentLen := spec[index]
		for k := s; k < segmentLen+s; k++ {
			out[k] = Empty
		}
	}

	return out
}

func ParseInput(line string) ([]byte, []int) {
	springs, groups, ok := strings.Cut(line, " ")
	if !ok {
		panic(errors.New("invalid format"))
	}
	intGroups := make([]int, 0)
	for _, s := range strings.Split(groups, ",") {
		intGroups = append(intGroups, lib.AsInt(s))
	}
	return []byte(springs), intGroups
}

func CountArrangements(input []string) (total int) {
	for _, line := range input {
		springs, groups := ParseInput(line)
		arrangements := FindArrangements(springs, groups)
		total += len(FilterValidSolutions(springs, groups, arrangements))
	}

	return
}
