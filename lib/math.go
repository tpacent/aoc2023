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
