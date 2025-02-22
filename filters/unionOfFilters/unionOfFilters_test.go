package unionOfFilters

import (
	"filters/allowed"
	blocked2 "filters/blocked"
	"filters/filter"
	"filters/filterByCriteria"
	"reflect"
	"testing"
)

func TestNewUnionOfFiltersInstance(t *testing.T) {
	instance := NewUnionOfFiltersInstance()

	if instance == nil {
		t.Fatal("Expected instance not to be nil")
	}
}

func TestUnionOfFilters_Add(t *testing.T) {
	allowedStrings := map[string]int{
		"grey":   1, // dummy values
		"apple":  2,
		"gopher": 3,
		"vrt":    4,
		"port":   5,
		"sigma":  12,
	}

	blockedStrings := map[string]int{
		"black": 2,
		"kusha": 3,
		"chica": 13,
		"kusa":  14,
		"admin": 4,
		"love":  6,
	}

	criteria1 := allowed.NewAllowedInstance(allowedStrings)
	criteria2 := blocked2.NewBlockedInstance(blockedStrings)
	filter1 := filterByCriteria.NewFilterByCriteria(criteria1)
	filter2 := filterByCriteria.NewFilterByCriteria(criteria2)

	filters := []filter.Filter{filter1, filter2}

	unionFilterInstance := NewUnionOfFiltersInstance()
	unionFilterInstance.Add(filter1)
	unionFilterInstance.Add(filter2)

	if !reflect.DeepEqual(filters, unionFilterInstance.Filters) {
		t.Errorf("Add function not adding filters properly : returned %v, expected %v", unionFilterInstance.Filters, filters)
	}
}

func TestUnionOfFilters_Accepts(t *testing.T) {
	allowedStrings := map[string]int{
		"grey":   1, // dummy values
		"apple":  2,
		"gopher": 3,
		"vrt":    4,
		"port":   5,
		"sigma":  12,
	}

	blockedStrings := map[string]int{
		"black": 2,
		"kusha": 3,
		"chica": 13,
		"kusa":  14,
		"admin": 4,
		"love":  6,
	}

	allowedCriteria := allowed.NewAllowedInstance(allowedStrings)
	blockedCriteria := blocked2.NewBlockedInstance(blockedStrings)

	filter1 := filterByCriteria.NewFilterByCriteria(allowedCriteria)
	filter2 := filterByCriteria.NewFilterByCriteria(blockedCriteria)

	unionInstance := NewUnionOfFiltersInstance()
	unionInstance.Add(filter1)
	unionInstance.Add(filter2)

	testCases := []struct {
		input    string
		expected bool
	}{
		{"kusha", false},
		{"chica", false},
		{"kusa", false},
		{"love", false},
		{"grey", true},
		{"port", true},
		{"sigma", true},
		{"", true},
	}

	for _, tc := range testCases {
		result := unionInstance.Accepts(tc.input)
		if tc.expected != result {
			t.Errorf("Accepts(%q) = %v expected %v", tc.input, result, tc.expected)
		}
	}

}
