package character

import (
	"github.com/yuriykis/pacman/move"
)

type MoverType int8

const (
	NoneMoverType MoverType = iota
	PlayerType
	GhostBlueType
	GhostRedType
	GhostWhiteType
	GhostOrangeType
)

type Mover interface {
	Move() move.Direction
}
