package day3_test

import (
	"aoc2023/day3"
	"aoc2023/lib"
	"testing"
)

const Day3Part1Solution = 528819

func TestSolveDay3Part1(t *testing.T) {
	data := lib.MustReadFile("testdata/input.txt")
	partNumbers := day3.FindPartNumbers(data)
	actual := lib.Sum(partNumbers...)
	if actual != Day3Part1Solution {
		t.Error("unexpected value")
	}
	t.Log(actual)
}

const Day3Part2Solution = 80403602

func TestSolveDay3Part2(t *testing.T) {
	data := lib.MustReadFile("testdata/input.txt")
	ratios := day3.FindGearsRatios(data)
	actual := lib.Sum(ratios...)
	if actual != Day3Part2Solution {
		t.Error("unexpected value")
	}
	t.Log(actual)
}

var testdata = []string{
	"467..114..",
	"...*......",
	"..35..633.",
	"......#...",
	"617*......",
	".....+.58.",
	"..592.....",
	"......755.",
	"...$.*....",
	".664.598..",
}

func TestFindPartNumbers(t *testing.T) {
	partNumbers := day3.FindPartNumbers(testdata)
	if len(partNumbers) != 8 {
		t.Error("unexpected part count")
	}
	actual := lib.Sum(partNumbers...)
	if actual != 4361 {
		t.Error("unexpected value")
	}
}

func TestFindGears(t *testing.T) {
	ratios := day3.FindGearsRatios(testdata)
	actual := lib.Sum(ratios...)
	if actual != 467835 {
		t.Error("unexpected value")
	}
}
