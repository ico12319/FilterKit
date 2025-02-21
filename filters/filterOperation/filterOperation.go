package filterOperation

import "filters/filter"

type FilterOperation interface {
	Add(f filter.Filter)
}
