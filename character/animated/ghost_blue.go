package animated

import (
	"math/rand"
	"pacman/move"
)

type GhostBlue struct {
	animated
}

func (char *GhostBlue) Move() move.Direction {
	return move.Direction(rand.Intn(4) + 1)
}
