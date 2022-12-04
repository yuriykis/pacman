package base

import (
	"pacman/board"
	"pacman/character/types"
	"pacman/utils"

	"fyne.io/fyne/v2"
)

type ICharacter interface {
	InitCharacter(pos *board.Position)
	CharacterImage() fyne.Resource
	CharacterType() types.CharacterType
	Position() *board.Position
	SetCharacterPosition(pos *board.Position)
	SetCharacterType(cType types.CharacterType)
}

type Character struct {
	img   fyne.Resource
	cType types.CharacterType
	pos   *board.Position
}

func (char *Character) CharacterImage() fyne.Resource {
	return char.img
}

func (char *Character) InitCharacter(pos *board.Position) {
	char.setCharacterImage()
	char.SetCharacterPosition(pos)
}

func (char *Character) CharacterType() types.CharacterType {
	return char.cType
}

func (char *Character) SetCharacterType(cType types.CharacterType) {
	char.cType = cType
}

func (char *Character) setCharacterImage() {
	char.img = utils.ResourceForCharacter(char.CharacterType())
}

func (char *Character) SetCharacterPosition(pos *board.Position) {
	char.pos = pos
	char.pos.SetFree(false)
}

func (char *Character) Position() *board.Position {
	return char.pos
}
