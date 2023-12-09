package day9_test

import (
	"aoc2023/day9"
	"aoc2023/lib"
	"testing"
)

const SolutionDay9Part1 = 2174807968

func TestSolveDay9Part1(t *testing.T) {
	var actual int
	for _, seq := range day9.ParseInput(lib.MustReadFile("testdata/input.txt")) {
		actual += day9.ExtrapolateTail(seq)
	}
	if actual != SolutionDay9Part1 {
		t.Error("unexpected value")
	}
	t.Log(actual)
}

const SolutionDay9Part2 = 1208

func TestSolveDay9Part2(t *testing.T) {
	var actual int
	for _, seq := range day9.ParseInput(lib.MustReadFile("testdata/input.txt")) {
		actual += day9.ExtrapolateHead(seq)
	}
	if actual != SolutionDay9Part2 {
		t.Error("unexpected value")
	}
	t.Log(actual)
}

var testdata = [][]int{
	{0, 3, 6, 9, 12, 15},
	{1, 3, 6, 10, 15, 21},
	{10, 13, 16, 21, 30, 45},
}
var testExpected = []int{18, 28, 68}

func TestExample(t *testing.T) {
	for index, seq := range testdata {
		if actual := day9.ExtrapolateTail(seq); actual != testExpected[index] {
			t.Error("unexpected value")
		}
	}
}
