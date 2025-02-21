package craftMaxLengthFilter

import (
	"filters/filter"
	"filters/filterByCriteria"
	"filters/lengthBlocker"
)

func CreateMaxLengthFilter(maxLenght int) filter.Filter {
	lengthFilter := lengthBlocker.NewLengthBlocker(maxLenght)
	toReturn := filterByCriteria.NewFilterByCriteria(lengthFilter)
	return toReturn
}
