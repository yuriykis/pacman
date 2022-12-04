package inanimate

import (
	"pacman/character/base"
	"pacman/character/types"
)

type IInanimate interface {
	base.ICharacter
}

type inanimate struct {
	base.Character
}

func NewInanimateCharacter(cType types.CharacterType) IInanimate {
	var newChar IInanimate
	switch cType {
	case types.TCoin:
		newChar = &Coin{}
	default:
		newChar = nil
	}
	if newChar != nil {
		newChar.SetCharacterType(cType)
	}
	return newChar
}
