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

var SolutionDay10Part2 = 451

func TestSolveDay10Part2(t *testing.T) {
	input := lib.MustReadFileBytes("testdata/input.txt")
	field := day10.ParseField(input)
	field.CleanTiles(field.WalkFrom(field.LocateStart()))
	// gotcha: it may be needed to replace S tile with a mathing pipe
	// to correctly solve all cases.
	actual := day10.CountEnclosed(field)
	if actual != SolutionDay10Part2 {
		t.Error("unexpected value")
	}
	t.Log(actual)
}

var testEnclosed = [][]byte{
	[]byte("FF7FSF7F7F7F7F7F---7"),
	[]byte("L|LJ||||||||||||F--J"),
	[]byte("FL-7LJLJ||||||LJL-77"),
	[]byte("F--JF--7||LJLJ7F7FJ-"),
	[]byte("L---JF-JLJ.||-FJLJJ7"),
	[]byte("|F|F-JF---7F7-L7L|7|"),
	[]byte("|FFJF7L7F-JF7|JL---7"),
	[]byte("7-L-JL7||F7|L7F-7F7|"),
	[]byte("L.L7LFJ|||||FJL7||LJ"),
	[]byte("L7JLJL-JLJLJL--JLJ.L"),
}

func TestEnclosedCount(t *testing.T) {
	field := day10.ParseField(testEnclosed)
	visited := field.WalkFrom(field.LocateStart())
	field.CleanTiles(visited)
	if actual := day10.CountEnclosed(field); actual != 10 {
		t.Log("unexpected value")
	}
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
