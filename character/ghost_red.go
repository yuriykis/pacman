package character

import (
	"pacman/utils"

	"fyne.io/fyne/v2"
)

type GhostRed struct {
	character
}

func (char *GhostRed) CharacterType() utils.CharacterType {
	return utils.TGhostRed
}

func (char *GhostRed) Move() {

}

func (char *GhostRed) InitCharacter() {
	char.setCharacterImage()
}

func (char *GhostRed) setCharacterImage() {
	img, ok := utils.ResourceForCharacter(char.CharacterType()).(*fyne.StaticResource)
	if ok {
		char.img = img
	}
}
