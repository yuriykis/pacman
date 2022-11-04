package main

type Board struct {
	positions []*Position
	boardView [][]BoardPositionType
}

type BoardPositionType string

const (
	WallChar  BoardPositionType = "#"
	SpaceChar BoardPositionType = " "
)

var boardView = [][]BoardPositionType{
	{WallChar, WallChar, WallChar, WallChar, WallChar, WallChar, WallChar, WallChar, WallChar, WallChar, WallChar, WallChar, WallChar, WallChar, WallChar, WallChar, WallChar, WallChar, WallChar, WallChar},
	{WallChar, SpaceChar, WallChar, SpaceChar, SpaceChar, WallChar, SpaceChar, SpaceChar, WallChar, SpaceChar, SpaceChar, SpaceChar, SpaceChar, WallChar, SpaceChar, SpaceChar, SpaceChar, SpaceChar, SpaceChar, WallChar},
	{WallChar, SpaceChar, WallChar, SpaceChar, SpaceChar, WallChar, SpaceChar, SpaceChar, WallChar, SpaceChar, SpaceChar, SpaceChar, SpaceChar, WallChar, SpaceChar, SpaceChar, SpaceChar, SpaceChar, SpaceChar, WallChar},
	{WallChar, SpaceChar, WallChar, SpaceChar, SpaceChar, SpaceChar, SpaceChar, SpaceChar, SpaceChar, SpaceChar, SpaceChar, SpaceChar, SpaceChar, SpaceChar, SpaceChar, SpaceChar, SpaceChar, SpaceChar, SpaceChar, WallChar},
	{WallChar, SpaceChar, WallChar, WallChar, WallChar, WallChar, SpaceChar, SpaceChar, WallChar, WallChar, WallChar, WallChar, WallChar, WallChar, WallChar, SpaceChar, SpaceChar, WallChar, WallChar, WallChar},
	{WallChar, SpaceChar, SpaceChar, SpaceChar, SpaceChar, SpaceChar, SpaceChar, SpaceChar, SpaceChar, SpaceChar, SpaceChar, SpaceChar, SpaceChar, WallChar, SpaceChar, SpaceChar, SpaceChar, SpaceChar, SpaceChar, WallChar},
	{WallChar, SpaceChar, SpaceChar, SpaceChar, SpaceChar, SpaceChar, SpaceChar, SpaceChar, SpaceChar, SpaceChar, SpaceChar, SpaceChar, SpaceChar, WallChar, SpaceChar, SpaceChar, SpaceChar, SpaceChar, SpaceChar, WallChar},
	{WallChar, SpaceChar, WallChar, WallChar, WallChar, WallChar, WallChar, SpaceChar, WallChar, SpaceChar, SpaceChar, SpaceChar, SpaceChar, WallChar, SpaceChar, SpaceChar, SpaceChar, SpaceChar, SpaceChar, WallChar},
	{WallChar, SpaceChar, WallChar, SpaceChar, SpaceChar, SpaceChar, WallChar, SpaceChar, WallChar, SpaceChar, SpaceChar, SpaceChar, SpaceChar, WallChar, SpaceChar, SpaceChar, SpaceChar, SpaceChar, SpaceChar, WallChar},
	{WallChar, SpaceChar, WallChar, SpaceChar, SpaceChar, SpaceChar, WallChar, SpaceChar, WallChar, SpaceChar, SpaceChar, SpaceChar, SpaceChar, SpaceChar, SpaceChar, SpaceChar, SpaceChar, SpaceChar, SpaceChar, WallChar},
	{WallChar, SpaceChar, WallChar, SpaceChar, SpaceChar, SpaceChar, WallChar, WallChar, WallChar, WallChar, WallChar, WallChar, WallChar, WallChar, SpaceChar, SpaceChar, SpaceChar, SpaceChar, SpaceChar, WallChar},
	{WallChar, SpaceChar, WallChar, SpaceChar, SpaceChar, SpaceChar, SpaceChar, SpaceChar, SpaceChar, SpaceChar, SpaceChar, SpaceChar, SpaceChar, WallChar, SpaceChar, SpaceChar, WallChar, WallChar, WallChar, WallChar},
	{WallChar, SpaceChar, WallChar, SpaceChar, SpaceChar, SpaceChar, WallChar, SpaceChar, SpaceChar, SpaceChar, SpaceChar, SpaceChar, SpaceChar, WallChar, SpaceChar, SpaceChar, SpaceChar, WallChar, SpaceChar, WallChar},
	{WallChar, SpaceChar, WallChar, SpaceChar, SpaceChar, SpaceChar, WallChar, SpaceChar, SpaceChar, SpaceChar, SpaceChar, SpaceChar, SpaceChar, WallChar, SpaceChar, SpaceChar, SpaceChar, WallChar, SpaceChar, WallChar},
	{WallChar, SpaceChar, WallChar, SpaceChar, WallChar, WallChar, WallChar, WallChar, WallChar, SpaceChar, SpaceChar, SpaceChar, SpaceChar, WallChar, SpaceChar, SpaceChar, SpaceChar, WallChar, SpaceChar, WallChar},
	{WallChar, SpaceChar, WallChar, SpaceChar, WallChar, SpaceChar, SpaceChar, SpaceChar, SpaceChar, SpaceChar, SpaceChar, SpaceChar, SpaceChar, WallChar, SpaceChar, SpaceChar, SpaceChar, WallChar, SpaceChar, WallChar},
	{WallChar, SpaceChar, WallChar, SpaceChar, WallChar, SpaceChar, SpaceChar, SpaceChar, SpaceChar, SpaceChar, SpaceChar, SpaceChar, SpaceChar, SpaceChar, SpaceChar, SpaceChar, SpaceChar, SpaceChar, SpaceChar, WallChar},
	{WallChar, SpaceChar, WallChar, SpaceChar, WallChar, SpaceChar, SpaceChar, SpaceChar, SpaceChar, SpaceChar, SpaceChar, SpaceChar, SpaceChar, SpaceChar, SpaceChar, SpaceChar, SpaceChar, WallChar, SpaceChar, WallChar},
	{WallChar, SpaceChar, WallChar, SpaceChar, WallChar, SpaceChar, SpaceChar, SpaceChar, SpaceChar, SpaceChar, SpaceChar, SpaceChar, SpaceChar, WallChar, SpaceChar, SpaceChar, SpaceChar, WallChar, SpaceChar, WallChar},
	{WallChar, WallChar, WallChar, WallChar, WallChar, WallChar, WallChar, WallChar, WallChar, WallChar, WallChar, WallChar, WallChar, WallChar, WallChar, WallChar, WallChar, WallChar, WallChar, WallChar},
}

func NewBoard() *Board {
	b := Board{}
	b.initBoardView()
	return &b
}

func (b *Board) initBoardView() {
	b.boardView = boardView
}

func (b *Board) BoardPositionTypeByPosition(pos Position) BoardPositionType {
	return b.boardView[pos.X][pos.Y]
}
