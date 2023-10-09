package character

import (
	"github.com/yuriykis/pacman/move"

	"fyne.io/fyne/v2"
)

const (
	PlayerHealth = 50
)

type Player struct {
	BaseCharacter

	currentDirection move.Direction
	score            int
	health           int
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

func (char *Player) Score() int {
	return char.score
}

func (char *Player) AddScore(c Collectible) {
	char.score += c.Value()
}

func (char *Player) Health() int {
	return char.health
}

func (char *Player) SetHealth(h int) {
	char.health = h
}
