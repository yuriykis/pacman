package types

type CharacterType int8

const (
	TNoAnimatedType CharacterType = iota
	TPlayer
	TGhostBlue
	TGhostRed
	TGhostWhite
	TGhostOrange
	TCoin
)
