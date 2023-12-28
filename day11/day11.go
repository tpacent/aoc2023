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
	return lib.Abs(a.X-b.X) + lib.Abs(a.Y-b.Y)
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

func IsEmpty[T comparable](row []T, empty T) bool {
	return slices.IndexFunc(row, func(value T) bool { return value != empty }) < 0
}

func SumExpandedDistances(coords []*Coords, empties *EmptyInfo, multiplier int) (total int) {
	for _, pair := range lib.PermuteMofN(coords, 2) {
		xMin, xMax := min(pair[0].X, pair[1].X), max(pair[0].X, pair[1].X)
		xEmpty := lib.CountMatches(empties.Cols[xMin:xMax], true)
		yMin, yMax := min(pair[0].Y, pair[1].Y), max(pair[0].Y, pair[1].Y)
		yEmpty := lib.CountMatches(empties.Rows[yMin:yMax], true)
		xDist := (lib.Abs(pair[0].X-pair[1].X) - xEmpty) + xEmpty*multiplier
		yDist := (lib.Abs(pair[0].Y-pair[1].Y) - yEmpty) + yEmpty*multiplier
		total += xDist + yDist
	}
	return
}

type EmptyInfo struct {
	Cols []bool
	Rows []bool
}

func GetEmptyInfo(data [][]byte, empty byte) *EmptyInfo {
	info := EmptyInfo{
		Rows: make([]bool, len(data)),
		Cols: make([]bool, len(data[0])),
	}

	for index, row := range data {
		if IsEmpty(row, empty) {
			info.Rows[index] = true
		}
	}

	for index, col := range Transpose(data) {
		if IsEmpty(col, empty) {
			info.Cols[index] = true
		}
	}

	return &info
}
