package animated

import (
	"pacman/character/base"
	"pacman/character/types"
	"pacman/move"
)

type IAnimated interface {
	Move() move.Direction
	base.ICharacter
}

type animated struct {
	base.Character
}

func NewAnimatedCharacter(cType types.CharacterType) IAnimated {
	var newChar IAnimated
	switch cType {
	case types.TPlayer:
		newChar = &Player{}
	case types.TGhostBlue:
		newChar = &GhostBlue{}
	case types.TGhostOrange:
		newChar = &GhostOrange{}
	case types.TGhostRed:
		newChar = &GhostRed{}
	case types.TGhostWhite:
		newChar = &GhostWhite{}
	default:
		newChar = nil
	}
	if newChar != nil {
		newChar.SetCharacterType(cType)
	}
	return newChar
}
