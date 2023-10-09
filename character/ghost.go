package character

import (
	"math/rand"

	"github.com/yuriykis/pacman/move"
)

type GhostBlue struct {
	BaseCharacter
	attacker
	mType MoverType
}

func (char *GhostBlue) Move() move.Direction {
	return move.Direction(rand.Intn(4) + 1)
}

type GhostOrange struct {
	BaseCharacter
	attacker
	mType MoverType
}

func (char *GhostOrange) Move() move.Direction {
	return move.Direction(rand.Intn(4) + 1)
}

type GhostRed struct {
	BaseCharacter
	attacker
	mType MoverType
}

func (char *GhostRed) Move() move.Direction {
	return move.Direction(rand.Intn(4) + 1)
}

type GhostWhite struct {
	BaseCharacter
	attacker
	mType MoverType
}

func (char *GhostWhite) Move() move.Direction {
	return move.Direction(rand.Intn(4) + 1)
}
