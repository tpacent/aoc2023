package day16_test

import (
	"aoc2023/day16"
	"aoc2023/lib"
	"testing"
)

const SolutionDay16Part1 = 7860

func TestSolveDay16Part1(t *testing.T) {
	input := lib.MustReadFileBytes("testdata/input.txt")
	visited := day16.TracePath(day16.CreateGrid(input), &day16.BeamState{Dir: day16.DirEast})
	actual := len(visited)
	if actual != SolutionDay16Part1 {
		t.Error("unexpected value")
	}
	t.Log(actual)
}

const SolutionDay16Part2 = 8331

func TestSolveDay16Part2(t *testing.T) {
	input := lib.MustReadFileBytes("testdata/input.txt")
	actual := day16.FindMaxVisits(input)
	if actual != SolutionDay16Part2 {
		t.Error("unexpected value")
	}
	t.Log(actual)
}

var testinput = [][]byte{
	[]byte(`.|...\....`),
	[]byte(`|.-.\.....`),
	[]byte(`.....|-...`),
	[]byte(`........|.`),
	[]byte(`..........`),
	[]byte(`.........\`),
	[]byte(`..../.\\..`),
	[]byte(`.-.-/..|..`),
	[]byte(`.|....-|.\`),
	[]byte(`..//.|....`),
}

func TestSample(t *testing.T) {
	visited := day16.TracePath(day16.CreateGrid(testinput), &day16.BeamState{Dir: day16.DirEast})
	if actual := len(visited); actual != 46 {
		t.Error("unexpected value", actual)
	}
}
