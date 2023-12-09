package day9

import (
	"aoc2023/lib"
	"strings"
)

func ExtrapolateTail(seq []int) int {
	layers := CreateLayers(seq)
	for k := len(layers) - 1; k > 0; k-- {
		layer := layers[k]
		upperLayer := layers[k-1]
		layers[k-1] = append(upperLayer, upperLayer[len(upperLayer)-1]+layer[len(layer)-1])
	}
	return layers[0][len(layers[0])-1]
}

func ExtrapolateHead(seq []int) int {
	layers := CreateLayers(seq)
	for k := len(layers) - 1; k > 0; k-- {
		layer := layers[k]
		upperLayer := layers[k-1]
		layers[k-1] = append([]int{upperLayer[0] - layer[0]}, layers[k-1]...)
	}
	return layers[0][0]
}

func CreateLayers(seq []int) [][]int {
	layers := [][]int{seq}
	for {
		prev := layers[len(layers)-1]
		layer := make([]int, 0)
		allZeroes := true

		for k := 1; k < len(prev); k++ {
			diff := prev[k] - prev[k-1]
			layer = append(layer, diff)
			if allZeroes && diff != 0 {
				allZeroes = false
			}
		}

		layers = append(layers, layer)

		if allZeroes {
			break
		}
	}
	return layers
}

func ParseInput(input []string) (seqs [][]int) {
	for _, line := range input {
		nums := strings.Fields(line)
		seq := make([]int, 0, len(nums))
		for _, n := range nums {
			seq = append(seq, lib.AsInt(n))
		}
		seqs = append(seqs, seq)
	}
	return
}
