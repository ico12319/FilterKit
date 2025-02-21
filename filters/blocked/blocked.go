package blocked

type Blocked struct {
	BlockedStrings map[string]int
}

func NewBlockedInstance(blockedStrings map[string]int) *Blocked {
	return &Blocked{BlockedStrings: blockedStrings}
}

func (b *Blocked) IsValid(s string) bool {
	_, ok := b.BlockedStrings[s]
	if ok {
		return false
	}
	return true
}
