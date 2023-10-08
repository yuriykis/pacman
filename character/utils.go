//go:generate fyne bundle -o bundled.go ../icons

package character

import (
	"fyne.io/fyne/v2"
)

func ResourceForMover(cType MoverType) fyne.Resource {
	switch cType {
	case PlayerType:
		return resourcePacmanSvg
	case GhostBlueType:
		return resourceGhostBlueSvg
	case GhostRedType:
		return resourceGhostRedSvg
	case GhostWhiteType:
		return resourceGhostWhiteSvg
	case GhostOrangeType:
		return resourceGhostOrangeSvg
	}
	return nil
}

func ResourceForCollectible(cType CollectibleType) fyne.Resource {
	switch cType {
	case CoinType:
		return resourceCoinSvg
	}
	return nil
}

func ResourceForWall() fyne.Resource {
	return resourceWallSvg
}
