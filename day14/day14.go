package day14

import (
	"aoc2023/lib"
	"hash/crc64"
	"slices"
)

const CubicStone byte = '#'
const RoundStone byte = 'O'
const EmptyTile byte = '.'

func RollStonesRight(data [][]byte) [][]byte {
	for _, row := range data {
		RollRow(row)
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

func FindAllIndex(vec []byte, needle byte) []int {
	indexes := make([]int, 0, 32)
	for index, value := range vec {
		if value == needle {
			indexes = append(indexes, index)
		}
	}
	return indexes
}

func CalcLoad(data [][]byte) (load int) {
	for _, row := range data {
		for index, value := range row {
			if value == RoundStone {
				load += index + 1
			}
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
	return loads[index+loopOffset]
}
