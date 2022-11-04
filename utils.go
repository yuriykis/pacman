//go:generate fyne bundle -o bundled.go icons

package main

import "fyne.io/fyne/v2"

func resourceForCharacter(cType CharacterType) fyne.Resource {
	switch cType {
	case Player:
		return resourcePacmanSvg
	case GhostBlue:
		return resourceGhostBlueSvg
	case GhostRed:
		return resourceGhostRedSvg
	case GhostWhite:
		return resourceGhostWhiteSvg
	case GhostOrange:
		return resourceGhostOrangeSvg
	}
	return nil
}

func resourceForWall() fyne.Resource {
	return resourceWallSvg
}
