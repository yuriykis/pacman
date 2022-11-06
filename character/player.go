package character

type Player struct {
	character
}

func (ghost *Player) Move() {

}

func (ghost *Player) SetCharacterImage() {
}

func (ghost *Player) CharacterImage() {

}

func (ghost *Player) CharacterType() CharacterType {
	return TPlayer
}
