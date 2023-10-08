package character

type CollectibleType int8

const (
	NoneCollectibleType CollectibleType = iota
	CoinType
)

type Collectible interface {
	Collect()
}

func CreateCollectible(cType CollectibleType) Collectible {
	var c Collectible
	switch cType {
	case CoinType:
		c = &Coin{}
	default:
		c = nil
	}
	// if c != nil {
	// 	c.SetCollectibleType(cType)
	// }
	return c
}
