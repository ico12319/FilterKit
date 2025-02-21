package filterByCriteria

import "filters/validator"

type FilterByCriteria[T validator.FailBasedFilter] struct {
	Criteria T
}

func NewFilterByCriteria[T validator.FailBasedFilter](criteria T) *FilterByCriteria[T] {
	return &FilterByCriteria[T]{Criteria: criteria}
}

func (fByCriteria *FilterByCriteria[T]) Accepts(s string) bool {
	return fByCriteria.Criteria.IsValid(s)
}
