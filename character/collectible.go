package character

type CollectibleType int8

const (
	NoneCollectibleType CollectibleType = iota
	CoinType
)

type Collectible interface {
	Collect()
}
