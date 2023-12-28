package day10_test

import (
	"aoc2023/day10"
	"aoc2023/lib"
	"testing"
)

var SolutionDay10Part1 = 6838

func TestSolveDay10Part1(t *testing.T) {
	input := lib.MustReadFileBytes("testdata/input.txt")
	field := day10.ParseField(input)
	actual := field.WalkFrom(field.LocateStart()).Size() / 2
	if actual != SolutionDay10Part1 {
		t.Error("unexpected value")
	}
	t.Log(actual)
}

var testinput = [][]byte{
	[]byte("-L|F7"),
	[]byte("7S-7|"),
	[]byte("L|7||"),
	[]byte("-L-J|"),
	[]byte("L|-JF"),
}

func TestParse(t *testing.T) {
	field := day10.ParseField(testinput)
	x, y := field.LocateStart()

	if x != 1 || y != 1 {
		t.Error("unexpected coordinates")
	}

	if field.SafeTile(3, 4) != day10.BendJ {
		t.Error("unexpected tile")
	}

	if field.SafeTile(-2, -8) != day10.None {
		t.Error("unexpected tile")
	}

	visitSet := field.WalkFrom(x, y)

	if visitSet.Size() != 8 {
		t.Error("unexpected steps")
	}
}
