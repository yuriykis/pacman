package character

type GhostRed struct {
	character
}

func (ghost *GhostRed) Move() {

}

func (ghost *GhostRed) SetCharacterImage() {
}

func (ghost *GhostRed) CharacterImage() {

}

func (ghost *GhostRed) CharacterType() CharacterType {
	return TGhostRed
}
