package board

import (
	"sync"

	"fyne.io/fyne/v2"
)

type PositionType int8

const (
	NoPositionType PositionType = iota
	Space
	Wall
)

type PositionData struct {
	X    int
	Y    int
	Cell fyne.CanvasObject
}
type Position struct {
	PositionData

	pType  PositionType
	isFree bool
	mu     sync.Mutex
}

func NewPosition(x int, y int, cell fyne.CanvasObject) *Position {
	pd := PositionData{
		X:    x,
		Y:    y,
		Cell: cell,
	}
	return &Position{
		PositionData: pd,
		pType:        Space,
		isFree:       true,
	}
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

func (pos *Position) IsFree() bool {
	return pos.isFree
}

func (pos *Position) SetFree(isFree bool) {
	pos.isFree = isFree
}

func (pos *Position) SetCell(cell fyne.CanvasObject) {
	pos.Cell = cell
}

func (pos *Position) GetCell() fyne.CanvasObject {
	return pos.Cell
}

func (pos *Position) Lock() {
	pos.mu.Lock()
}

func (pos *Position) Unlock() {
	pos.mu.Unlock()
}
