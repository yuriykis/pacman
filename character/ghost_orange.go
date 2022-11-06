package character

type GhostOrange struct {
	character
}

func (ghost *GhostOrange) Move() {

}

func (ghost *GhostOrange) SetCharacterImage() {
}

func (ghost *GhostOrange) CharacterImage() {

}

func (ghost *GhostOrange) CharacterType() CharacterType {
	return TGhostOrange
}
