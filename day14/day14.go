package day14

import (
	"bytes"
	"slices"
)

const CubicStone byte = '#'
const RoundStone byte = 'O'
const EmptyTile byte = '.'

func RollStonesRight(data [][]byte) [][]byte {
	for index := range data {
		RollRow(data[index])
	}
	return data
}

func RollRow(vec []byte) {
	stones := FindAllIndex(vec, RoundStone)
	slices.Reverse(stones)

	for _, stonepos := range stones {
		runlen := RunLength(vec[stonepos+1:], EmptyTile)
		nextPos := stonepos + runlen
		vec[stonepos] = EmptyTile
		vec[nextPos] = RoundStone
	}
}

func RunLength(vec []byte, needle byte) (runlen int) {
	for _, curr := range vec {
		if curr != needle {
			break
		}
		runlen++
	}
	return
}

func FindAllIndex(vec []byte, needle byte) (indexes []int) {
	cursor := 0

	for {
		if cursor >= len(vec) {
			break
		}
		index := bytes.IndexByte(vec[cursor:], needle)
		if index < 0 {
			break
		}
		indexes = append(indexes, cursor+index)
		cursor += index + 1
	}

	return
}

func CalcLoad(data [][]byte) (load int) {
	for _, row := range data {
		for _, index := range FindAllIndex(row, RoundStone) {
			load += index + 1
		}
	}

	return
}
