package types

type CharacterType int8

const (
	TNoCharacterType CharacterType = iota
	TPlayer
	TGhostBlue
	TGhostRed
	TGhostWhite
	TGhostOrange
	TCoin
)
