package character

import (
	"pacman/move"

	"fyne.io/fyne/v2"
)

type Player struct {
	character
	currentDirection move.Direction
}

func (char *Player) Move() move.Direction {
	return char.currentDirection
}

func (char *Player) SetDirection(dir move.Direction) {
	char.currentDirection = dir
}

func (char *Player) MapKeyCodeToDirection(k fyne.KeyEvent) {
	switch k.Name {
	case fyne.KeyUp:
		char.SetDirection(move.Up)
	case fyne.KeyRight:
		char.SetDirection(move.Right)
	case fyne.KeyDown:
		char.SetDirection(move.Down)
	case fyne.KeyLeft:
		char.SetDirection(move.Left)
	}
}
