package character

import (
	"math/rand"
	"pacman/move"
)

type Player struct {
	character
}

func (char *Player) Move() move.Direction {
	return move.Direction(rand.Intn(4))
}
