//go:generate fyne bundle -o bundled.go ../icons

package character

import (
	"fyne.io/fyne/v2"
)

func ResourceForMover(mType MoverType) fyne.Resource {
	switch mType {
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

func ResourceForCollectible(mType CollectibleType) fyne.Resource {
	switch mType {
	case CoinType:
		return resourceCoinSvg
	}
	return nil
}

func ResourceForWall() fyne.Resource {
	return resourceWallSvg
}
