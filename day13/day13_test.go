package day13_test

import (
	"aoc2023/day13"
	"aoc2023/lib"
	"bytes"
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
	for k := 1; k < len(hInput); k++ {
		if CheckReflect(hInput, k-1, k) {
			t.Log("Found", k)
		}
	}

}

func CheckReflect(input [][]byte, upper, lower int) bool {
	for {
		if upper < 0 || lower > len(input)-1 {
			break
		}

		if !bytes.Equal(input[upper], input[lower]) {
			return false
		}

		upper--
		lower++
	}

	return true
}
