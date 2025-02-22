package intersectionOfFilters

import (
	"filters/allowed"
	"filters/blocked"
	"filters/filter"
	"filters/filterByCriteria"
	"reflect"
	"testing"
)

func TestNewIntersectionOfFiltersInstance(t *testing.T) {
	instance := NewIntersectionOfFiltersInstance()

	if instance == nil {
		t.Fatal("Expected instance not to be nil")
	}
}

func TestIntersectionOfFilters_Add(t *testing.T) {
	allowedStrings := map[string]int{
		"eyes": 1,
		"eva":  3,
	}
	blockedStrings := map[string]int{
		"ratatui": 1,
		"kisa":    2,
	}

	allowedInstance := allowed.NewAllowedInstance(allowedStrings)
	blockedInstance := blocked.NewBlockedInstance(blockedStrings)

	filter1 := filterByCriteria.NewFilterByCriteria(allowedInstance)
	filter2 := filterByCriteria.NewFilterByCriteria(blockedInstance)

	filters := []filter.Filter{filter1, filter2}

	intersectionInstance := NewIntersectionOfFiltersInstance()
	intersectionInstance.Add(filter1)
	intersectionInstance.Add(filter2)

	if !reflect.DeepEqual(filters, intersectionInstance.Filters) {
		t.Errorf("Add function not adding filters properly : returned %v, expected %v", intersectionInstance.Filters, filters)
	}
}

func TestIntersectionOfFilters_Accepts(t *testing.T) {
	allowedStrings := map[string]int{
		"okay":   1, // dummy values
		"no":     2,
		"bussin": 3,
		"kicin":  4,
	}

	blockedStrings := map[string]int{
		"grey": 1,
		"girl": 2,
		"boy":  3,
		"okay": 4,
	}

	testCases := []struct {
		input    string
		expected bool
	}{
		{"okay", false},
		{"grey", false},
		{"girl", false},
		{"no", true},
		{"kicin", true},
	}

	allowedInstance := allowed.NewAllowedInstance(allowedStrings)
	blockedInstance := blocked.NewBlockedInstance(blockedStrings)

	filter1 := filterByCriteria.NewFilterByCriteria(allowedInstance)
	filter2 := filterByCriteria.NewFilterByCriteria(blockedInstance)

	intersectionInstance := NewIntersectionOfFiltersInstance()
	intersectionInstance.Add(filter1)
	intersectionInstance.Add(filter2)

	for _, tc := range testCases {
		result := intersectionInstance.Accepts(tc.input)
		if result != tc.expected {
			t.Errorf("Accepts(%q) = %v expected %v", tc.input, result, tc.expected)
		}
	}
}
