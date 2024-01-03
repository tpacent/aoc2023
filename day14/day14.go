package day14

import (
	"aoc2023/lib"
	"bytes"
	"hash/crc64"
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

func HashData(data [][]byte) uint64 {
	checksum := crc64.New(crc64.MakeTable(crc64.ECMA))
	for _, row := range data {
		if _, err := checksum.Write(row); err != nil {
			panic(err)
		}
	}
	return checksum.Sum64()
}

func LoadAtIteration(data [][]byte, iterations int) int {
	hashIndexes := make(map[uint64]int)
	var loopOffset int
	var loopSize int
	var loads = make([]int, 0)

	for k := 0; k < iterations; k++ {
		loads = append(loads, CalcLoad(data))

		key := HashData(data)
		if index, ok := hashIndexes[key]; ok {
			loopOffset = index
			loopSize = k - index
			break
		} else {
			hashIndexes[key] = k
		}

		data = RollStonesRight(data)
		data = lib.RotateCW(data)
		data = RollStonesRight(data)
		data = lib.RotateCW(data)
		data = RollStonesRight(data)
		data = lib.RotateCW(data)
		data = RollStonesRight(data)
		data = lib.RotateCW(data)
	}

	index := (iterations - loopOffset) % loopSize
	return loads[loopOffset:][index]
}
