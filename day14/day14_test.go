package day14_test

import (
	"aoc2023/day14"
	"aoc2023/lib"
	"slices"
	"testing"
)

const SolutionDay14Part1 = 109755

func TestSolveDay14Part1(t *testing.T) {
	data := lib.MustReadFileBytes("testdata/input.txt")
	actual := day14.CalcLoad(day14.RollStonesRight(lib.RotateCW(data)))
	if actual != SolutionDay14Part1 {
		t.Error("unexpected value")
	}
	t.Log(actual)
}

func TestFindAllIndex(t *testing.T) {
	data := []byte("00...0....0..00")
	expected := []int{0, 1, 5, 10, 13, 14}
	if actual := day14.FindAllIndex(data, 'O'); slices.Equal(actual, expected) {
		t.Log("unexpected data")
	}
}

var sampledata = [][]byte{
	[]byte("O....#...."),
	[]byte("O.OO#....#"),
	[]byte(".....##..."),
	[]byte("OO.#O....O"),
	[]byte(".O.....O#."),
	[]byte("O.#..O.#.#"),
	[]byte("..O..#O..O"),
	[]byte(".......O.."),
	[]byte("#....###.."),
	[]byte("#OO..#...."),
}

func TestRoll(t *testing.T) {
	actual := day14.RollStonesRight(lib.RotateCW(sampledata))
	t.Log(actual)
}

func TestRotateClockwise(t *testing.T) {
	actual := lib.RotateCW(sampledata)
	t.Log(actual)
}

func TestLoad(t *testing.T) {
	rolled := day14.RollStonesRight(lib.RotateCW(sampledata))
	if actual := day14.CalcLoad(rolled); actual != 136 {
		t.Log("unexpected value")
	}
}
