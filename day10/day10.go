package day10

import (
	"aoc2023/lib"
	"bytes"
)

const (
	PipeV byte = '|'
	PipeH byte = '-'
	BendL byte = 'L'
	BendJ byte = 'J'
	Bend7 byte = '7'
	BendF byte = 'F'
	None  byte = '.'
	Start byte = 'S'
)

type Dir uint8

const (
	DirUp = 1 << iota
	DirDown
	DirLeft
	DirRight
	DirNone = 0
)

var moveDirections = []Dir{DirUp, DirRight, DirDown, DirLeft}

var dirInverse = map[Dir]Dir{
	DirDown:  DirUp,
	DirLeft:  DirRight,
	DirUp:    DirDown,
	DirRight: DirLeft,
}

var connectors = map[byte]Dir{
	None:  DirNone,
	Start: DirUp | DirDown | DirLeft | DirRight,
	PipeV: DirUp | DirDown,
	PipeH: DirLeft | DirRight,
	Bend7: DirLeft | DirDown,
	BendJ: DirUp | DirLeft,
	BendF: DirDown | DirRight,
	BendL: DirUp | DirRight,
}

func canMove(from, onto byte, dir Dir) bool {
	if from == 0 || onto == 0 {
		return false
	}

	if connectors[from]&dir == 0 {
		return false
	}

	if connectors[onto]&dirInverse[dir] == 0 {
		return false
	}

	return true
}

type Landscape struct {
	Data   []byte
	Width  int
	Height int
}

func (l *Landscape) LocateStart() (x, y int) {
	index := bytes.IndexByte(l.Data, Start)
	row := index / l.Width
	col := index % l.Width
	return col, row
}

// SafeTile will not panic on out of bounds coords
func (l *Landscape) SafeTile(x, y int) byte {
	index := l.TileIndex(x, y)
	if index < 0 || index > len(l.Data)-1 {
		return None
	}
	return l.Data[index]
}

func (l *Landscape) TileIndex(x, y int) int {
	return y*l.Width + x
}

func (l *Landscape) PeekTile(x, y int, dir Dir) (int, int, byte) {
	switch dir {
	case DirLeft:
		x--
	case DirRight:
		x++
	case DirDown:
		y++
	case DirUp:
		y--
	}
	return x, y, l.SafeTile(x, y)
}

func (l *Landscape) WalkFrom(x, y int) lib.Set[int] {
	visited := make(lib.Set[int])
	currentTile := l.SafeTile(x, y)

	for {
		moved := false
		visited.Add(l.TileIndex(x, y))

		for _, dir := range moveDirections {
			toX, toY, toTile := l.PeekTile(x, y, dir)

			// check visited
			if visited.Has(l.TileIndex(toX, toY)) {
				continue
			}

			// check can move in dir
			if !canMove(currentTile, toTile, dir) {
				continue
			}

			moved = true
			x, y = toX, toY
			currentTile = toTile
			break
		}

		if !moved {
			break
		}
	}

	return visited
}

func ParseField(lines [][]byte) *Landscape {
	w := len(lines[0])
	h := len(lines)
	buf := bytes.NewBuffer(make([]byte, 0, w*h))
	for _, line := range lines {
		_, err := buf.Write(line)
		if err != nil {
			panic(err)
		}
	}
	return &Landscape{
		Width:  w,
		Height: h,
		Data:   buf.Bytes(),
	}
}
