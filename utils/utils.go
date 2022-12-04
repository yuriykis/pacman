//go:generate fyne bundle -o bundled.go ../icons

package utils

import (
	"pacman/character/types"

	"fyne.io/fyne/v2"
)

func ResourceForCharacter(cType types.CharacterType) fyne.Resource {
	switch cType {
	case types.TPlayer:
		return resourcePacmanSvg
	case types.TGhostBlue:
		return resourceGhostBlueSvg
	case types.TGhostRed:
		return resourceGhostRedSvg
	case types.TGhostWhite:
		return resourceGhostWhiteSvg
	case types.TGhostOrange:
		return resourceGhostOrangeSvg
	}
	return nil
}

func ResourceForWall() fyne.Resource {
	return resourceWallSvg
}

func ResourceForItem(iType types.CharacterType) fyne.Resource {
	switch iType {
	case types.TCoin:
		return resourceCoinSvg
	}
	return nil
}
