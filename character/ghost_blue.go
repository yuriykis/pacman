package character

import (
	"math/rand"
	"pacman/move"
)

type GhostBlue struct {
	character
}

func (char *GhostBlue) Move() move.Direction {
	return move.Direction(rand.Intn(4))
}
