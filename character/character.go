package character

import (
	"pacman/board"
	"pacman/utils"

	"fyne.io/fyne/v2"
)

type ICharacter interface {
	Move()
	InitCharacter(pos *board.Position)
	CharacterImage() *fyne.StaticResource
	CharacterType() utils.CharacterType
	Position() *board.Position
	SetCharacterPosition(pos *board.Position)
}

type character struct {
	img   *fyne.StaticResource
	cType utils.CharacterType
	pos   *board.Position
}

func (char *character) CharacterImage() *fyne.StaticResource {
	return char.img
}

func (char *character) InitCharacter(pos *board.Position) {
	char.setCharacterImage()
	char.SetCharacterPosition(pos)
}

func (char *character) CharacterType() utils.CharacterType {
	return char.cType
}

func (char *character) setCharacterImage() {
	img, ok := utils.ResourceForCharacter(char.CharacterType()).(*fyne.StaticResource)
	if ok {
		char.img = img
	}
}

func (char *character) SetCharacterPosition(pos *board.Position) {
	char.pos = pos
}

func (char *character) Position() *board.Position {
	return char.pos
}

func NewCharacter(cType utils.CharacterType) ICharacter {
	switch cType {
	case utils.TPlayer:
		return &Player{character: character{cType: utils.TPlayer}}
	case utils.TGhostBlue:
		return &GhostBlue{character: character{cType: utils.TGhostBlue}}
	case utils.TGhostOrange:
		return &GhostOrange{character: character{cType: utils.TGhostOrange}}
	case utils.TGhostRed:
		return &GhostRed{character: character{cType: utils.TGhostRed}}
	case utils.TGhostWhite:
		return &GhostWhite{character: character{cType: utils.TGhostWhite}}
	}
	return nil
}
