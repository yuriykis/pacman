package character

type Coin struct {
	BaseCharacter
	cType CollectibleType
}

func (char *Coin) Collect() {
	return
}
