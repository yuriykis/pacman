package character

import (
	"pacman/utils"

	"fyne.io/fyne/v2"
)

type ICharacter interface {
	Move()
	InitCharacter()
	CharacterImage() *fyne.StaticResource
	CharacterType() utils.CharacterType
}

type character struct {
	img   *fyne.StaticResource
	cType utils.CharacterType
}

func (char *character) CharacterImage() *fyne.StaticResource {
	return char.img
}

func (char *character) InitCharacter() {
	char.setCharacterImage()
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
