package lib

func PermuteMofN[T any](items []T, m int) [][]T {
	out := make([][]T, 0)

	if m == 1 {
		for _, item := range items {
			out = append(out, []T{item})
		}
		return out
	}

	for k := 0; k < len(items)-m+1; k++ {
		item := items[k]
		rest := PermuteMofN(items[k+1:], m-1)

		for _, tail := range rest {
			out = append(out, append([]T{item}, tail...))
		}
	}

	return out
}
