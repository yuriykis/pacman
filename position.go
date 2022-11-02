package main

type PositionType int8

const (
	NoPositionType PositionType = iota
	FreePosition
	Wall
)

type Position struct {
	X      int
	Y      int
	pType  PositionType
	isFree bool
}

func newPosition(x int, y int, pType PositionType, isFree bool) *Position {
	return &Position{X: x, Y: y, isFree: isFree}
}
