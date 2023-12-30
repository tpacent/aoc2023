package day11_test

import (
	"aoc2023/day11"
	"aoc2023/lib"
	"testing"
)

const SolutionDay11Part1 = 9329143

func TestSolveDay11Part1(t *testing.T) {
	data := lib.MustReadFileBytes("testdata/input.txt")
	actual := day11.SumExpandedDistances(
		day11.CollectItems(data, '#'),
		day11.GetEmptyInfo(data, '.'),
		2,
	)
	if actual != SolutionDay11Part1 {
		t.Error("unexpected value")
	}
	t.Log(actual)
}

const SolutionDay11Part2 = 710674907809

func TestSolveDay11Part2(t *testing.T) {
	data := lib.MustReadFileBytes("testdata/input.txt")
	actual := day11.SumExpandedDistances(
		day11.CollectItems(data, '#'),
		day11.GetEmptyInfo(data, '.'),
		1_000_000,
	)
	if actual != SolutionDay11Part2 {
		t.Error("unexpected value")
	}
	t.Log(actual)
}

var testdata = [][]byte{
	{1, 2, 0, 4},
	{0, 0, 0, 0},
	{9, 10, 0, 12},
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
	empties := day11.GetEmptyInfo(sample, '.')
	coords := day11.CollectItems(sample, '#')
	if actual := day11.SumExpandedDistances(coords, empties, 2); actual != 374 {
		t.Error("unexpected value")
	}
}

func TestEmpty(t *testing.T) {
	empties := day11.GetEmptyInfo(sample, '.')
	coords := day11.CollectItems(sample, '#')
	actual := day11.SumExpandedDistances(coords, empties, 10)
	t.Log(actual)
}
