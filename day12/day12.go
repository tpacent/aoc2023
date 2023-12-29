package day12

import (
	"aoc2023/lib"
	"bytes"
	"errors"
	"fmt"
	"strings"
)

const (
	Empty   = '.'
	Damaged = '#'
	Unknown = '?'
)

func CacheKey(data []byte, spec []int) string {
	return string(data) + fmt.Sprintf("%v", spec)
}

func FindArrangements(data []byte, spec []int, offset int, cache map[string][][]int) (indexes [][]int) {
	if idxs, ok := cache[CacheKey(data, spec)]; ok {
		return AddOffsetToIndexes(idxs, offset)
	}

	segmentLen, rest := spec[0], spec[1:]

	for k := 0; k < len(data)-segmentLen+1; k++ {
		if !checkSegment(data, k, segmentLen) {
			continue
		}

		if len(rest) == 0 {
			indexes = append(indexes, []int{k + offset})
			continue
		}

		nextSliceIndex := k + segmentLen + 1
		if nextSliceIndex >= len(data) {
			continue
		}

		for _, sub := range FindArrangements(data[nextSliceIndex:], rest, offset+nextSliceIndex, cache) {
			rec := append([]int{k + offset}, sub...)
			indexes = append(indexes, rec)
		}
	}

	indexes = FilterValidSolutions(data, spec, AddOffsetToIndexes(indexes, -offset))

	if len(spec) > 1 {
		cache[CacheKey(data, spec)] = indexes
	}

	return AddOffsetToIndexes(indexes, offset)
}

func AddOffsetToIndexes(indexes [][]int, offset int) [][]int {
	if len(indexes) == 0 {
		return indexes
	}

	out := make([][]int, 0, len(indexes))
	for _, idx := range indexes {
		row := make([]int, 0, len(idx))
		for _, n := range idx {
			row = append(row, n+offset)
		}
		out = append(out, row)
	}
	return out
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
		if IsValidSolution(ApplySolution(data, spec, s)) {
			valids = append(valids, s)
		}
	}
	return valids
}

func IsValidSolution(data []byte) bool {
	return bytes.IndexByte(data, Damaged) < 0
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

func CountArrangements(input []string, repeat int) (total int) {
	for _, line := range input {
		springs, groups := ParseInput(line)
		springs = lib.Repeat(springs, repeat)
		groups = lib.Repeat(groups, repeat)
		arrangements := FindArrangements(springs, groups, 0, make(map[string][][]int))
		total += len(arrangements)
	}

	return
}
