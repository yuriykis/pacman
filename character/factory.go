package character

import "math/rand"

func NewCharacter(mType CharacterType) BaseCharacter {
	var (
		mt = mType.MoverType
		bc = &baseCharacter{CType: mType}
	)

	if mt != NoneMoverType {
		switch mt {
		case PlayerType:
			return &Player{
				BaseCharacter: bc,
				score:         0,
				health:        PlayerHealth,
			}
		case GhostBlueType:
			return &GhostBlue{
				BaseCharacter: bc,
				attacker:      attacker{aType: AttackerType(rand.Intn(2) + 1)},
			}
		case GhostOrangeType:
			return &GhostOrange{
				BaseCharacter: bc,
				attacker:      attacker{aType: AttackerType(rand.Intn(2) + 1)},
			}
		case GhostRedType:
			return &GhostRed{
				BaseCharacter: bc,
				attacker:      attacker{aType: AttackerType(rand.Intn(2) + 1)},
			}
		case GhostWhiteType:
			return &GhostWhite{
				BaseCharacter: bc,
				attacker:      attacker{aType: AttackerType(rand.Intn(2) + 1)},
			}
		default:
			return nil
		}
	}
	ct := mType.CollectibleType
	if ct != NoneCollectibleType {
		switch mType.CollectibleType {
		case CoinType:
			return &Coin{
				BaseCharacter: bc,
				value:         1,
			}
		default:
			return nil
		}
	}
	return nil
}
