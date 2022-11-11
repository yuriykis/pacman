package character

import (
	"math/rand"
	"pacman/move"
)

type GhostRed struct {
	character
}

func (char *GhostRed) Move() move.Direction {
	return move.Direction(rand.Intn(4))
}
