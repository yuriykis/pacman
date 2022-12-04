package inanimate

import (
	"pacman/character"
	"pacman/character/types"
)

type IInanimate interface {
	character.ICharacter
}

type inanimate struct {
	character.Character
}

func NewAnimatedCharacter(cType types.CharacterType) IInanimate {
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
