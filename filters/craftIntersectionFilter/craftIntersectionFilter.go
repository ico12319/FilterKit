package craftIntersectionFilter

import (
	"filters/craftAllowedFilter"
	"filters/craftBlockedFilter"
	"filters/craftMaxLengthFilter"
	"filters/filter"
	"filters/intersectionOfFilters"
)

func CreateIntersectionFilter(fileNameAllowed string, fileNameBlocked string, maxLength int) filter.Filter {
	lengthFilter := craftMaxLengthFilter.CreateMaxLengthFilter(maxLength)
	allowedFilter := craftAllowedFilter.CreateAllowedFilter(fileNameAllowed)
	blockedFilter := craftBlockedFilter.CreateBlockedFilter(fileNameBlocked)

	intersection := intersectionOfFilters.NewIntersectionOfFiltersInstance()
	intersection.Add(lengthFilter)
	intersection.Add(allowedFilter)
	intersection.Add(blockedFilter)
	return intersection
}
