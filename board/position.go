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
	isFree bool
}

func NewPosition(x int, y int) *Position {
	return &Position{X: x, Y: y, pType: Space, isFree: true}
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
