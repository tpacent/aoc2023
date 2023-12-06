package day6_test

import (
	"aoc2023/day6"
	"aoc2023/lib"
	"testing"
)

var Day6Part1Solution = 3317888

func TestSolveDay6Part1(t *testing.T) {
	races := day6.ParseRaces(lib.MustReadFile("testdata/input.txt"))
	actual := 1
	for _, race := range races {
		actual *= day6.CountWaysToWin(race)
	}
	if actual != Day6Part1Solution {
		t.Error("unexpected value")
	}
	t.Log(actual)
}

var Day6Part2Solution = 24655068

func TestSolveDay6Part2(t *testing.T) {
	race := day6.ParseLongRace(lib.MustReadFile("testdata/input.txt"))
	actual := day6.CountWaysToWin(race)
	if actual != Day6Part2Solution {
		t.Error("unexpected value")
	}
	t.Log(actual)
}

var testinput = []string{
	"Time:      7  15   30",
	"Distance:  9  40  200",
}

func TestCount(t *testing.T) {
	races := day6.ParseRaces(testinput)
	if actual := day6.CountWaysToWin(races[0]); actual != 4 {
		t.Error("unexpected value")
	}
	if actual := day6.CountWaysToWin(races[1]); actual != 8 {
		t.Error("unexpected value")
	}
	if actual := day6.CountWaysToWin(races[2]); actual != 9 {
		t.Error("unexpected value")
	}
}
