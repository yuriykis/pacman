package character

import (
	"pacman/character/animated"
	"pacman/character/base"
	"pacman/character/inanimate"
	"pacman/character/types"
)

func NewCharacter(cType types.CharacterType) base.ICharacter {
	switch cType {
	case types.TPlayer, types.TGhostBlue, types.TGhostOrange, types.TGhostRed, types.TGhostWhite:
		return animated.NewAnimatedCharacter(cType)
	case types.TCoin:
		return inanimate.NewInanimateCharacter(cType)
	default:
		return nil
	}
}
