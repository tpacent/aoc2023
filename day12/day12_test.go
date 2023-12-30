package day12_test

import (
	"aoc2023/day12"
	"aoc2023/lib"
	"testing"
)

const SolutionDay12Part1 = 6935

func TestDay12Part1(t *testing.T) {
	actual := day12.CountArrangements(lib.MustReadFile("testdata/input.txt"), 1)
	if actual != SolutionDay12Part1 {
		t.Error("unexpected value")
	}
	t.Log(actual)
}

const SolutionDay12Part2 = 3920437278260

func TestDay12Part2(t *testing.T) {
	actual := day12.CountArrangements(lib.MustReadFile("testdata/input.txt"), 5)
	if actual != SolutionDay12Part2 {
		t.Error("unexpected value")
	}
	t.Log(actual)
}

type TestCase struct {
	Springs  []byte
	Groups   []int
	Expected int
}

var cases = []*TestCase{
	{Springs: []byte("???.###"), Groups: []int{1, 1, 3}, Expected: 1},
	{Springs: []byte(".??..??...?##."), Groups: []int{1, 1, 3}, Expected: 4},
	{Springs: []byte("?#?#?#?#?#?#?#?"), Groups: []int{1, 3, 1, 6}, Expected: 1},
	{Springs: []byte("????.#...#..."), Groups: []int{4, 1, 1}, Expected: 1},
	{Springs: []byte("????.######..#####."), Groups: []int{1, 6, 5}, Expected: 4},
	{Springs: []byte("?###????????"), Groups: []int{3, 2, 1}, Expected: 10},
}

func TestNumArrangements(t *testing.T) {
	sum := 0

	for _, testcase := range cases {
		actual := day12.NumArrangements(testcase.Springs, testcase.Groups)
		if actual != testcase.Expected {
			t.Error("unexpected value", actual)
		}
		sum += actual
	}

	if sum != 21 {
		t.Error("unexpected sum", sum)
	}
}

func TestUnfold(t *testing.T) {
	springs, groups := day12.Unfold([]byte("??????.??..?"), []int{2, 1, 2}, 5)
	actual := day12.NumArrangements(springs, groups)
	t.Log(actual)
}
