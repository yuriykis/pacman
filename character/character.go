package character

import "fyne.io/fyne/v2/canvas"

type ICharacter interface {
	Move()
	SetCharacterImage()
	CharacterImage()
	CharacterType() CharacterType
}

type character struct {
	img *canvas.Image
}

type CharacterType int8

const (
	TPlayer CharacterType = iota
	TGhostBlue
	TGhostRed
	TGhostWhite
	TGhostOrange
)

func NewCharacter(cType CharacterType) ICharacter {
	switch cType {
	case TPlayer:
		return &Player{}
	case TGhostBlue:
		return &GhostBlue{}
	case TGhostOrange:
		return &GhostOrange{}
	case TGhostRed:
		return &GhostRed{}
	case TGhostWhite:
		return &GhostWhite{}
	}
	return nil
}
