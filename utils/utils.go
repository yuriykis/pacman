//go:generate fyne bundle -o bundled.go ../icons

package utils

import (
	"pacman/character"

	"fyne.io/fyne/v2"
)

func ResourceForCharacter(cType character.CharacterType) fyne.Resource {
	switch cType {
	case character.TPlayer:
		return resourcePacmanSvg
	case character.TGhostBlue:
		return resourceGhostBlueSvg
	case character.TGhostRed:
		return resourceGhostRedSvg
	case character.TGhostWhite:
		return resourceGhostWhiteSvg
	case character.TGhostOrange:
		return resourceGhostOrangeSvg
	}
	return nil
}

func ResourceForWall() fyne.Resource {
	return resourceWallSvg
}
