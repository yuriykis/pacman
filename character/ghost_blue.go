package character

import (
	"pacman/utils"

	"fyne.io/fyne/v2"
)

type GhostBlue struct {
	character
}

func (char *GhostBlue) CharacterType() utils.CharacterType {
	return utils.TGhostBlue
}

func (char *GhostBlue) Move() {
}

func (char *GhostBlue) InitCharacter() {
	char.setCharacterImage()
}

func (char *GhostBlue) setCharacterImage() {
	img, ok := utils.ResourceForCharacter(char.CharacterType()).(*fyne.StaticResource)
	if ok {
		char.img = img
	}
}
