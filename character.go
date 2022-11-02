//go:generate fyne bundle -o bundled.go characters

package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type Character struct {
	pos   *Position
	img   *canvas.Image
	cType CharacterType
}

type CharacterType int8

const (
	Player CharacterType = iota
	EnemyBlue
	EnemyYellow
	EnemyRed
)

func (char *Character) Position() *Position {
	return char.pos
}

func (char *Character) SetPosition(x int, y int) {
	char.pos.X = x
	char.pos.Y = y
}

func newCharacter(pos *Position, img *canvas.Image, cType CharacterType) *Character {
	return &Character{pos: pos, img: img, cType: cType}
}

func resourceForCharacter(isUser bool) fyne.Resource {
	if isUser {
		return resourcePacmanIconSvg
	}
	return nil
}
