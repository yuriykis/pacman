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

type Position struct {
	X      int
	Y      int
	pType  PositionType
	isFree bool
	cell   fyne.CanvasObject
	mu     sync.Mutex
}

func NewPosition(x int, y int, cell fyne.CanvasObject) *Position {
	return &Position{X: x, Y: y, pType: Space, isFree: true, cell: cell}
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
	pos.cell = cell
}

func (pos *Position) Cell() fyne.CanvasObject {
	return pos.cell
}

func (pos *Position) Lock() {
	pos.mu.Lock()
}

func (pos *Position) Unlock() {
	pos.mu.Unlock()
}
