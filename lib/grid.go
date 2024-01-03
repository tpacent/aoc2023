package lib

import "fmt"

func NewGrid[T any](dimensionSizes ...int) *Grid[T] {
	var zero T
	var size = Mul(dimensionSizes...)

	return &Grid[T]{
		data: make([]T, size),
		dims: dimensionSizes,
		zero: zero,
	}
}

// Grid is a generic container for n-dimensional grids
type Grid[T any] struct {
	data []T
	dims []int
	zero T // cached helper value to return in case of error
}

func (grid *Grid[T]) Len() int {
	return len(grid.data)
}

// Get returns a value at specified coordinates
func (grid *Grid[T]) Get(coords ...int) (T, error) {
	index, err := grid.Index(coords...)
	if err != nil {
		return grid.zero, err
	}
	return grid.data[index], nil
}

// Set updates a value at specified coordinates
func (grid *Grid[T]) Set(value T, coords ...int) error {
	index, err := grid.Index(coords...)
	if err != nil {
		return err
	}
	grid.data[index] = value
	return nil
}

// Index converts coordinates x, y, z, ... into a number
func (grid *Grid[T]) Index(coords ...int) (index int, err error) {
	if err = grid.validateCoords(coords...); err != nil {
		return
	}
	planeSize := 1
	for dIndex, dValue := range coords {
		index += dValue * planeSize
		planeSize *= grid.dims[dIndex]
	}
	return
}

// Coords converts index to a corresponding set of coords
func (grid *Grid[T]) Coords(index int) ([]int, error) {
	if err := grid.validateIndex(index); err != nil {
		return nil, err
	}
	coords := make([]int, len(grid.dims))
	planeSize := len(grid.data)
	// most significant coordinate last,
	// as in: x, y, z.
	for k := len(grid.dims) - 1; k > -1; k-- {
		planeSize /= grid.dims[k]
		coords[k] = index / planeSize
		index -= coords[k] * planeSize
	}
	return coords, nil
}

// validateIndex index validation routines
func (grid *Grid[T]) validateIndex(index int) error {
	if index < 0 {
		return fmt.Errorf("negative index not allowed")
	}
	if index < len(grid.data) {
		return nil
	}
	return fmt.Errorf("index too large")
}

// validateCoords coordinate set validation routines
func (grid *Grid[T]) validateCoords(coords ...int) error {
	if len(coords) != len(grid.dims) {
		return fmt.Errorf("unexpected coords len")
	}
	for index, c := range coords {
		if c < 0 || c >= grid.dims[index] {
			return fmt.Errorf("coord %d out of range", index)
		}
	}
	return nil
}
