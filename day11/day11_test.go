package day11_test

import (
	"aoc2023/day11"
	"aoc2023/lib"
	"testing"
)

var SolutionDay11Part1 = 9329143

func TestSolveDay11Part1(t *testing.T) {
	data := lib.MustReadFileBytes("testdata/input.txt")
	actual := day11.CountDistancePairs(data)
	if actual != SolutionDay11Part1 {
		t.Error("unexpected value")
	}
	t.Log(actual)
}

var testdata = [][]byte{
	{1, 2, 0, 4},
	{0, 0, 0, 0},
	{9, 10, 0, 12},
}

func TestTranspose(t *testing.T) {
	actual := day11.Transpose(testdata)

	if len(actual) != 4 {
		t.Error("unexpected slice len")
	}

	if actual[3][2] != 12 {
		t.Error("unexpected value")
	}
}

func TestExpand(t *testing.T) {
	actual := day11.ExpandSpace(testdata, 0)
	if len(actual) != 4 || len(actual[0]) != 5 {
		t.Error("unexpected expand")
	}
}

var sample = [][]byte{
	[]byte("...#......"),
	[]byte(".......#.."),
	[]byte("#........."),
	[]byte(".........."),
	[]byte("......#..."),
	[]byte(".#........"),
	[]byte(".........#"),
	[]byte(".........."),
	[]byte(".......#.."),
	[]byte("#...#....."),
}

func TestExample(t *testing.T) {
	data := make([][]byte, len(sample))

	for index := range data {
		data[index] = sample[index]
	}

	if actual := day11.CountDistancePairs(data); actual != 374 {
		t.Error("unexpected value")
	}
}
