package character

type Coin struct {
	BaseCharacter

	mType CollectibleType
	value int
}

func (char *Coin) Collect() {
	return
}

func (char *Coin) Value() int {
	return char.value
}
