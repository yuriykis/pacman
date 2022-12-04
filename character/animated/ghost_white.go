package animated

import (
	"math/rand"
	"pacman/move"
)

type GhostWhite struct {
	animated
}

func (char *GhostWhite) Move() move.Direction {
	return move.Direction(rand.Intn(4) + 1)
}
