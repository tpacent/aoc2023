package day16

import (
	"aoc2023/lib"
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

type VisitMap map[uint16]map[uint8]struct{}

func (vm VisitMap) Visit(beam BeamState) (visited bool) {
	key := KeyFunc(beam.X, beam.Y)
	dirstore, ok := vm[key]
	if ok {
		// check if dir was already visited
		if _, visited := dirstore[beam.Dir]; visited {
			return true
		}
		dirstore[beam.Dir] = struct{}{}
		return
	}
	store := make(map[uint8]struct{}, 4)
	store[beam.Dir] = struct{}{}
	vm[key] = store
	return
}

// KeyFunc works for coordinates up to 256
func KeyFunc(x, y int) uint16 {
	return uint16(x)<<8 | uint16(y)
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
			len(TracePath(grid, BeamState{X: x, Y: 0, Dir: DirSouth})),
			len(TracePath(grid, BeamState{X: x, Y: h - 1, Dir: DirNorth})),
		)
	}
	for y := 0; y < h; y++ {
		maxVisits = max(
			maxVisits,
			len(TracePath(grid, BeamState{X: 0, Y: y, Dir: DirEast})),
			len(TracePath(grid, BeamState{X: w - 1, Y: y, Dir: DirWest})),
		)
	}
	return
}

func TracePath(grid *lib.Grid[byte], seed BeamState) VisitMap {
	visited := make(VisitMap, grid.Len())
	beams := []BeamState{seed}

	for {
		newStates := make([]BeamState, 0)

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

func ProcessTile(beam BeamState, tile byte) []BeamState {
	nextDirs, ok := nextDirections[beam.Dir][tile]
	if !ok {
		return []BeamState{beam} // keep going
	}
	states := make([]BeamState, 0, 2)
	for _, next := range nextDirs {
		states = append(states, BeamState{Dir: next, X: beam.X, Y: beam.Y})
	}
	return states
}

func NextState(beam BeamState) BeamState {
	switch beam.Dir {
	case DirNorth:
		beam.Y--
	case DirSouth:
		beam.Y++
	case DirWest:
		beam.X--
	case DirEast:
		beam.X++
	}
	return beam
}
