package character

import (
	"math/rand"

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

type mover struct {
	mType MoverType
}

func (m *mover) Move() move.Direction {
	switch m.mType {
	case GhostBlueType:
		return move.Direction(rand.Intn(4) + 1)
	case GhostRedType:
		return move.Direction(rand.Intn(4) + 1)
	case GhostWhiteType:
		return move.Direction(rand.Intn(4) + 1)
	case GhostOrangeType:
		return move.Direction(rand.Intn(4) + 1)
	default:
		return move.Direction(rand.Intn(4) + 1)
	}
}
