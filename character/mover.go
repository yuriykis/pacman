package character

import (
	"pacman/move"
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

func CreateMover(cType MoverType) Mover {
	var newChar Mover
	switch cType {
	case PlayerType:
		newChar = &Player{}
	case GhostBlueType:
		newChar = &GhostBlue{}
	case GhostOrangeType:
		newChar = &GhostOrange{}
	case GhostRedType:
		newChar = &GhostRed{}
	case GhostWhiteType:
		newChar = &GhostWhite{}
	default:
		newChar = nil
	}
	// if newChar != nil {
	// 	newChar.SetType(cType)
	// }
	return newChar
}
