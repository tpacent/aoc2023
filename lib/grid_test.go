package lib_test

import (
	"aoc2023/lib"
	"slices"
	"testing"
)

func TestGrid(t *testing.T) {
	grid := lib.NewGrid[int]([]int{3, 3, 3}...)

	{
		actual, err := grid.Index(2, 1, 2)

		if err != nil {
			t.Fatal(err)
		}

		if actual != 23 {
			t.Error("unexpected value", 23)
		}
	}

	{
		actual, err := grid.Coords(23)

		if err != nil {
			t.Fatal(err)
		}

		if !slices.Equal(actual, []int{2, 1, 2}) {
			t.Error("unexpected value", actual)
		}
	}
}
