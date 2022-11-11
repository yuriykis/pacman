package board

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
	IsFree bool
}

func NewPosition(x int, y int) *Position {
	return &Position{X: x, Y: y, pType: Space, IsFree: true}
}

func (pos *Position) SetPositionType(pType PositionType) {
	if pType == Wall {
		pos.IsFree = false
	}
	pos.pType = pType
}

func (pos *Position) PositionType() PositionType {
	return pos.pType
}
