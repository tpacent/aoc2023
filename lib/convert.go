package lib

import "strconv"

func AsInt(s string) int {
	if n, err := strconv.Atoi(s); err == nil {
		return n
	}

	panic("unexpected input")
}
