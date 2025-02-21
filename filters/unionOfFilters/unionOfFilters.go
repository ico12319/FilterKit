package unionOfFilters

import "filters/filter"

type UnionOfFilters struct {
	Filters []filter.Filter
}

func NewUnionOfFiltersInstance() *UnionOfFilters {
	return &UnionOfFilters{}
}

func (u *UnionOfFilters) Add(f filter.Filter) {
	u.Filters = append(u.Filters, f)
}
func (u *UnionOfFilters) Accepts(s string) bool {
	for i := 0; i < len(u.Filters); i++ {
		if u.Filters[i].Accepts(s) {
			return true
		}
	}
	return false
}
