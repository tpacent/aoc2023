package day11

import (
	"aoc2023/lib"
	"slices"
)

type Coords struct {
	X int
	Y int
}

func Distance(a, b *Coords) int {
	return lib.Mod(a.X-b.X) + lib.Mod(a.Y-b.Y)
}

func CountDistancePairs(data [][]byte) (dists int) {
	data = ExpandSpace(data, '.')
	coords := CollectItems(data, '#')
	for _, pair := range lib.PermuteMofN(coords, 2) {
		dists += Distance(pair[0], pair[1])
	}
	return
}

func CollectItems[T comparable](data [][]T, item T) (coords []*Coords) {
	h := len(data)
	w := len(data[0])

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if data[y][x] == item {
				coords = append(coords, &Coords{X: x, Y: y})
			}
		}
	}

	return
}

func Transpose[T any](data [][]T) [][]T {
	w := len(data[0])
	h := len(data)

	out := make([][]T, w)

	for k := 0; k < len(out); k++ {
		out[k] = make([]T, h)
	}

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			out[x][y] = data[y][x]
		}
	}

	return out
}

func ExpandRows[T comparable](data [][]T, empty T) (out [][]T) {
	for _, row := range data {
		out = append(out, row)

		if slices.IndexFunc(row, func(value T) bool { return value != empty }) < 0 {
			out = append(out, row) // expand space
		}

	}

	return
}

func ExpandSpace[T comparable](data [][]T, empty T) [][]T {
	data = ExpandRows(data, empty)
	data = Transpose(data)
	data = ExpandRows(data, empty)
	data = Transpose(data)
	return data
}
