package character

type GhostWhite struct {
	character
}

func (ghost *GhostWhite) Move() {

}

func (ghost *GhostWhite) SetCharacterImage() {
}

func (ghost *GhostWhite) CharacterImage() {

}

func (ghost *GhostWhite) CharacterType() CharacterType {
	return TGhostWhite
}
