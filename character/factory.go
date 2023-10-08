package character

func NewCharacter(cType CharacterType) BaseCharacter {
	var (
		mt = cType.MoverType
		bc = &baseCharacter{CType: cType}
	)

	if mt != NoneMoverType {
		switch mt {
		case PlayerType:
			return &Player{
				BaseCharacter: bc,
			}
		case GhostBlueType:
			return &GhostBlue{
				BaseCharacter: bc,
			}
		case GhostOrangeType:
			return &GhostOrange{
				BaseCharacter: bc,
			}
		case GhostRedType:
			return &GhostRed{
				BaseCharacter: bc,
			}
		case GhostWhiteType:
			return &GhostWhite{
				BaseCharacter: bc,
			}
		default:
			return nil
		}
	}
	ct := cType.CollectibleType
	if ct != NoneCollectibleType {
		switch cType.CollectibleType {
		case CoinType:
			return &Coin{
				BaseCharacter: bc,
			}
		default:
			return nil
		}
	}
	return nil
}
