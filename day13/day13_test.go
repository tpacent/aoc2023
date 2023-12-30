package day13_test

import (
	"aoc2023/day13"
	"aoc2023/lib"
	"testing"
)

const SolutionDay13Part1 = 36015

func TestSolveDay13Part1(t *testing.T) {
	actual := day13.ParseInput(lib.MustReadFileBytes("testdata/input.txt"))
	if actual != SolutionDay13Part1 {
		t.Error("unexpected value")
	}
	t.Log(actual)
}

var hInput = [][]byte{
	[]byte("#...##..#"),
	[]byte("#....#..#"),
	[]byte("..##..###"),
	[]byte("#####.##."),
	[]byte("#####.##."),
	[]byte("..##..###"),
	[]byte("#....#..#"),
}

func TestData(t *testing.T) {
	actual := day13.FindReflection(hInput)
	if actual != 4 {
		t.Error("unexpected value")
	}
}
