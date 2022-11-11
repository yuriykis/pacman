package character

import (
	"math/rand"
	"pacman/move"
)

type GhostOrange struct {
	character
}

func (char *GhostOrange) Move() move.Direction {
	return move.Direction(rand.Intn(4))
}
