package character

type AttackerType int8

const (
	NoneAttackerType AttackerType = iota
	LowAttackerType
	MediumAttackerType
	HighAttackerType
)

type Attacker interface {
	Attack(int) int
}

type attacker struct {
	aType AttackerType
}

func (a *attacker) Attack(pHealth int) int {
	r := pHealth - int(a.aType)
	if r < 0 {
		return 0
	}
	return r
}
