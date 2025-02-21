package intersectionOfFilters

import "filters/filter"

type IntersectionOfFilters struct {
	Filters []filter.Filter
}

func NewIntersectionOfFiltersInstance() *IntersectionOfFilters {
	return &IntersectionOfFilters{}
}

func (inter *IntersectionOfFilters) Add(f filter.Filter) {
	inter.Filters = append(inter.Filters, f)
}

func (inter *IntersectionOfFilters) Accepts(s string) bool {
	for i := 0; i < len(inter.Filters); i++ {
		if !inter.Filters[i].Accepts(s) {
			return false
		}
	}
	return true
}
