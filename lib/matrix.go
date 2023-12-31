package lib

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

func RotateCW[T any](data [][]T) [][]T {
	w := len(data[0])
	h := len(data)

	out := make([][]T, w)
	for k := 0; k < w; k++ {
		out[k] = make([]T, h)
	}

	for y, row := range data {
		outY := h - y - 1
		for x, value := range row {
			out[x][outY] = value
		}
	}

	return out
}
