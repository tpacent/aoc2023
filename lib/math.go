package lib

import (
	"golang.org/x/exp/constraints"
)

func Sum[T constraints.Integer | constraints.Float](numbers ...T) (sum T) {
	for _, n := range numbers {
		sum += n
	}

	return
}

func Mul[T constraints.Integer | constraints.Float](numbers ...T) (product T) {
	product = 1

	for _, n := range numbers {
		product *= n
	}

	return
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func Mod[T constraints.Signed | constraints.Float](value T) T {
	if value < 0 {
		return -1 * value
	}
	return value
}
