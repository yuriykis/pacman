package character

import (
	"pacman/utils"

	"fyne.io/fyne/v2"
)

type Player struct {
	character
}

func (char *Player) CharacterType() utils.CharacterType {
	return utils.TPlayer
}

func (char *Player) InitCharacter() {
	char.setCharacterImage()
}

func (char *Player) setCharacterImage() {
	img, ok := utils.ResourceForCharacter(char.CharacterType()).(*fyne.StaticResource)
	if ok {
		char.img = img
	}
}

func (char *Player) Move() {

}
