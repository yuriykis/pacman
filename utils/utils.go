//go:generate fyne bundle -o bundled.go ../icons

package utils

import (
	"fyne.io/fyne/v2"
)

func ResourceForCharacter(cType CharacterType) fyne.Resource {
	switch cType {
	case TPlayer:
		return resourcePacmanSvg
	case TGhostBlue:
		return resourceGhostBlueSvg
	case TGhostRed:
		return resourceGhostRedSvg
	case TGhostWhite:
		return resourceGhostWhiteSvg
	case TGhostOrange:
		return resourceGhostOrangeSvg
	}
	return nil
}

func ResourceForWall() fyne.Resource {
	return resourceWallSvg
}

func ResourceForItem(iType ItemType) fyne.Resource {
	switch iType {
	case TCoin:
		return resourceCoinSvg
	}
	return nil
}
