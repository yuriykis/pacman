//go:generate fyne bundle -o bundled.go characters

package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type Character struct {
	img   *canvas.Image
	cType CharacterType
}

type CharacterType int8

const (
	Player CharacterType = iota
	GhostBlue
	GhostRed
	GhostWhite
)

func newCharacter(img *canvas.Image, cType CharacterType) *Character {
	return &Character{img: img, cType: cType}
}

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
	}
	return nil
}
