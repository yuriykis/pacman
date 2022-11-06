package character

import (
	"pacman/utils"

	"fyne.io/fyne/v2"
)

type GhostWhite struct {
	character
}

func (char *GhostWhite) CharacterType() utils.CharacterType {
	return utils.TGhostWhite
}

func (char *GhostWhite) Move() {

}

func (char *GhostWhite) InitCharacter() {
	char.setCharacterImage()
}

func (char *GhostWhite) setCharacterImage() {
	img, ok := utils.ResourceForCharacter(char.CharacterType()).(*fyne.StaticResource)
	if ok {
		char.img = img
	}
}
