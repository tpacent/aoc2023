package lib

func CountMatches[T comparable](items []T, needle T) (count int) {
	for _, item := range items {
		if item == needle {
			count++
		}
	}

	return
}
