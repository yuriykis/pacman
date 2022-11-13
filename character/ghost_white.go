package character

import (
	"math/rand"
	"pacman/move"
)

type GhostWhite struct {
	character
}

func (char *GhostWhite) Move() move.Direction {
	return move.Direction(rand.Intn(4) + 1)
}
