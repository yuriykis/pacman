package board

import (
	"math/rand"

	"fyne.io/fyne/v2"
)

type Board struct {
	positions []*Position
	boardView [][]BoardPositionType
}

type BoardPositionType string

const (
	WallChar  BoardPositionType = "#"
	SpaceChar BoardPositionType = " "
)

func NewBoard() *Board {
	b := Board{}
	b.initBoardView()
	return &b
}

func (b *Board) initBoardView() {
	b.boardView = boardView
}

func (b *Board) SetPositionTypes() {
	for _, p := range b.positions {
		if b.BoardPositionTypeByPosition(*&p.PositionData) == WallChar {
			p.SetPositionType(Wall)
			continue
		}
		if b.BoardPositionTypeByPosition(*&p.PositionData) == SpaceChar {
			p.SetPositionType(Space)
			continue
		}
	}
}

func (b *Board) BoardPositionTypeByPosition(pos PositionData) BoardPositionType {
	return b.boardView[pos.X][pos.Y]
}

func (b *Board) FindFreePosition() *Position {
	var freePositions []*Position
	for _, p := range b.positions {
		if p.IsFree() {
			freePositions = append(freePositions, p)
		}
	}
	if len(freePositions) != 0 {
		return freePositions[rand.Intn(len(freePositions))]
	}
	return nil
}

func (b *Board) CreatePosition(x int, y int, cell fyne.CanvasObject) PositionData {
	pos := NewPosition(x, y, cell)
	b.positions = append(b.positions, pos)
	return pos.PositionData
}

func (b *Board) FindPosition(x int, y int) *Position {
	for _, p := range b.positions {
		if p.X == x && p.Y == y {
			return p
		}
	}
	return nil
}

func (b *Board) Positions() []*Position {
	return b.positions
}
