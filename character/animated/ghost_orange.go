package animated

import (
	"math/rand"
	"pacman/move"
)

type GhostOrange struct {
	animated
}

func (char *GhostOrange) Move() move.Direction {
	return move.Direction(rand.Intn(4) + 1)
}
