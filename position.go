package main

import "errors"

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

func (pos *Position) CharacterFromPosition() (*Character, error) {
	char := pos.char
	if char == nil {
		return nil, errors.New("no character on given position")
	}
	return char, nil
}

func (pos *Position) SetPositionType(pType PositionType) {
	if pType == Wall {
		pos.isFree = false
	}
	pos.pType = pType
}

func (pos *Position) PositionType() PositionType {
	return pos.pType
}
