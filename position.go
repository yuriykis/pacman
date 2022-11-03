package main

type PositionType int8

const (
	NoPositionType PositionType = iota
	Space
	Wall
)

type Position struct {
	X      int
	Y      int
	pType  PositionType
	isFree bool

	char *Character
}

func newPosition(x int, y int) *Position {
	return &Position{X: x, Y: y, pType: Space, isFree: true, char: nil}
}

func (pos *Position) assignCharacterToPosition(char *Character) {
	pos.char = char
	pos.isFree = false
}
