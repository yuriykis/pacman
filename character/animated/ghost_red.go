package animated

import (
	"math/rand"
	"pacman/move"
)

type GhostRed struct {
	animated
}

func (char *GhostRed) Move() move.Direction {
	return move.Direction(rand.Intn(4) + 1)
}
