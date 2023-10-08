package character

import (
	"math/rand"
	"pacman/move"
)

type GhostBlue struct {
	BaseCharacter

	cType MoverType
}

func (char *GhostBlue) Move() move.Direction {
	return move.Direction(rand.Intn(4) + 1)
}

type GhostOrange struct {
	BaseCharacter

	cType MoverType
}

func (char *GhostOrange) Move() move.Direction {
	return move.Direction(rand.Intn(4) + 1)
}

type GhostRed struct {
	BaseCharacter

	cType MoverType
}

func (char *GhostRed) Move() move.Direction {
	return move.Direction(rand.Intn(4) + 1)
}

type GhostWhite struct {
	BaseCharacter

	cType MoverType
}

func (char *GhostWhite) Move() move.Direction {
	return move.Direction(rand.Intn(4) + 1)
}
