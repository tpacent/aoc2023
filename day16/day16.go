package day16

import (
	"aoc2023/lib"
	"encoding/binary"
)

const (
	DirUnspecified uint8 = iota
	DirNorth
	DirWest
	DirSouth
	DirEast
)

type BeamState struct {
	X   int
	Y   int
	Dir uint8
}

type VisitMap map[string]map[uint8]struct{}

func (vm VisitMap) Visit(beam *BeamState) (visited bool) {
	buf := make([]byte, 0, 20)
	buf = binary.AppendUvarint(buf, uint64(beam.X))
	buf = binary.AppendUvarint(buf, uint64(beam.Y))
	dirstore, ok := vm[string(buf)]

	if ok {
		// check if dir was already visited
		if _, visited := dirstore[beam.Dir]; visited {
			return true
		}
		dirstore[beam.Dir] = struct{}{}
		return
	}

	vm[string(buf)] = map[uint8]struct{}{beam.Dir: {}}
	return
}

func CreateGrid(input [][]byte) *lib.Grid[byte] {
	grid := lib.NewGrid[byte](len(input[0]), len(input))
	for y, row := range input {
		for x, value := range row {
			if err := grid.Set(value, x, y); err != nil {
				panic(err)
			}
		}
	}
	return grid
}

func FindMaxVisits(input [][]byte) (maxVisits int) {
	w := len(input[0])
	h := len(input)

	grid := CreateGrid(input)

	for x := 0; x < w; x++ {
		maxVisits = max(
			maxVisits,
			len(TracePath(grid, &BeamState{X: x, Y: 0, Dir: DirSouth})),
			len(TracePath(grid, &BeamState{X: x, Y: h - 1, Dir: DirNorth})),
		)
	}

	for y := 0; y < h; y++ {
		maxVisits = max(
			maxVisits,
			len(TracePath(grid, &BeamState{X: 0, Y: y, Dir: DirEast})),
			len(TracePath(grid, &BeamState{X: w - 1, Y: y, Dir: DirWest})),
		)
	}

	return
}

func TracePath(grid *lib.Grid[byte], seed *BeamState) VisitMap {
	visited := make(VisitMap)
	beams := []*BeamState{seed}

	for {
		newStates := make([]*BeamState, 0)

		for _, beam := range beams {
			tile, err := grid.Get(beam.X, beam.Y)

			if err != nil {
				continue // terminate: out of range
			}

			// tile ok, set visited
			if ok := visited.Visit(beam); ok {
				continue // terminate: loop detected
			}

			for _, state := range ProcessTile(beam, tile) {
				newStates = append(newStates, NextState(state))
			}
		}

		if len(newStates) == 0 {
			break
		}

		beams = newStates
	}

	return visited
}

var nextDirections = map[uint8]map[byte][]uint8{
	DirEast: {
		'|':  {DirNorth, DirSouth},
		'\\': {DirSouth},
		'/':  {DirNorth},
	},
	DirWest: {
		'|':  {DirNorth, DirSouth},
		'\\': {DirNorth},
		'/':  {DirSouth},
	},
	DirNorth: {
		'-':  {DirWest, DirEast},
		'\\': {DirWest},
		'/':  {DirEast},
	},
	DirSouth: {
		'-':  {DirWest, DirEast},
		'\\': {DirEast},
		'/':  {DirWest},
	},
}

func ProcessTile(beam *BeamState, tile byte) (states []*BeamState) {
	nextDirs, ok := nextDirections[beam.Dir][tile]

	if !ok {
		return []*BeamState{beam} // keep going
	}

	for _, next := range nextDirs {
		states = append(states, &BeamState{Dir: next, X: beam.X, Y: beam.Y})
	}

	return
}

func NextState(beam *BeamState) *BeamState {
	state := BeamState{Dir: beam.Dir, X: beam.X, Y: beam.Y}
	switch beam.Dir {
	case DirNorth:
		state.Y--
	case DirSouth:
		state.Y++
	case DirWest:
		state.X--
	case DirEast:
		state.X++
	}
	return &state
}
