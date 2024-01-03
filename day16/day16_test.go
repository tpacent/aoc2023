package day16_test

import (
	"aoc2023/day16"
	"aoc2023/lib"
	"testing"
)

const SolutionDay16Part1 = 7860

func TestSolveDay16Part1(t *testing.T) {
	visited := day16.TracePath(lib.MustReadFileBytes("testdata/input.txt"))
	if actual := len(visited); actual != SolutionDay16Part1 {
		t.Error("unexpected value")
	}
	t.Log(len(visited))
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
	visited := day16.TracePath(testinput)
	if actual := len(visited); actual != 46 {
		t.Error("unexpected value", actual)
	}
}
