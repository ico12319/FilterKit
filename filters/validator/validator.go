package validator

type FailBasedFilter interface {
	IsValid(s string) bool
}
