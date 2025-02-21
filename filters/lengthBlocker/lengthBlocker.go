package lengthBlocker

import "unicode/utf8"

type LengthBlocker struct {
	MaxLength int
}

func NewLengthBlocker(maxLength int) *LengthBlocker {
	return &LengthBlocker{MaxLength: maxLength}
}

func (l *LengthBlocker) IsValid(s string) bool {
	return utf8.RuneCountInString(s) <= l.MaxLength
}
