package filter

type Filter interface {
	Accepts(s string) bool
}
