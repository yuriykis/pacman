package character

import (
	"pacman/utils"

	"fyne.io/fyne/v2"
)

type GhostOrange struct {
	character
}

func (char *GhostOrange) CharacterType() utils.CharacterType {
	return utils.TGhostOrange
}

func (char *GhostOrange) Move() {
}

func (char *GhostOrange) InitCharacter() {
	char.setCharacterImage()
}

func (char *GhostOrange) setCharacterImage() {
	img, ok := utils.ResourceForCharacter(char.CharacterType()).(*fyne.StaticResource)
	if ok {
		char.img = img
	}
}
