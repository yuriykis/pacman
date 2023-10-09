package character

import (
	"github.com/yuriykis/pacman/board"

	"fyne.io/fyne/v2"
)

type BaseCharacter interface {
	CharacterImage() fyne.Resource
	Type() CharacterType
	SetType(CharacterType)
	InitCharacter(*board.Position)
	SetCharacterPosition(*board.Position)
	Position() *board.Position
}

type CharacterType struct {
	MoverType       MoverType
	CollectibleType CollectibleType
}

type baseCharacter struct {
	img   fyne.Resource
	pos   *board.Position
	CType CharacterType
}

func (char *baseCharacter) CharacterImage() fyne.Resource {
	return char.img
}

func (char *baseCharacter) Type() CharacterType {
	return char.CType
}

func (char *baseCharacter) SetType(cType CharacterType) {
	char.CType = cType
}

func (char *baseCharacter) InitCharacter(pos *board.Position) {
	char.setCharacterImage()
	char.SetCharacterPosition(pos)
}

func (char *baseCharacter) setCharacterImage() {
	if char.CType.MoverType != NoneMoverType {
		char.img = ResourceForMover(char.Type().MoverType)
		return
	}
	// TODO: character probably should be either a mover or a collectible
	if char.CType.CollectibleType != NoneCollectibleType {
		char.img = ResourceForCollectible(char.Type().CollectibleType)
	}
}

func (char *baseCharacter) SetCharacterPosition(pos *board.Position) {
	char.pos = pos
	char.pos.SetFree(false)
}

func (char *baseCharacter) Position() *board.Position {
	return char.pos
}
