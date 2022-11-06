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
	img *fyne.StaticResource
}

func (char *character) CharacterImage() *fyne.StaticResource {
	return char.img
}

// func (char *character) InitCharacter() {
// 	char.setCharacterImage()
// }

// func (char *character) CharacterType() utils.CharacterType {
// 	return utils.TNoCharacterType
// }

// func (char *character) setCharacterImage() {
// 	img, ok := utils.ResourceForCharacter(char.CharacterType()).(*fyne.StaticResource)
// 	if ok {
// 		char.img = img
// 	}
// }

func NewCharacter(cType utils.CharacterType) ICharacter {
	switch cType {
	case utils.TPlayer:
		return &Player{}
	case utils.TGhostBlue:
		return &GhostBlue{}
	case utils.TGhostOrange:
		return &GhostOrange{}
	case utils.TGhostRed:
		return &GhostRed{}
	case utils.TGhostWhite:
		return &GhostWhite{}
	}
	return nil
}
