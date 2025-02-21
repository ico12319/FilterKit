package allowed

type Allowed struct {
	AllowedStrings map[string]int
}

func NewAllowedInstance(allowedStrings map[string]int) *Allowed {
	return &Allowed{AllowedStrings: allowedStrings}
}

func (a *Allowed) IsValid(s string) bool {
	_, ok := a.AllowedStrings[s]
	if ok {
		return true
	}
	return false
}
