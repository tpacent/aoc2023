package lib

func CountMatches[T comparable](items []T, needle T) (count int) {
	for _, item := range items {
		if item == needle {
			count++
		}
	}

	return
}

func Repeat[T any](src []T, count int) []T {
	dst := make([]T, len(src)*count)
	for k := 0; k < count; k++ {
		copy(dst[k*len(src):], src)
	}
	return dst
}
