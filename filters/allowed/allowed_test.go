package allowed

import (
	"reflect"
	"testing"
)

func TestNewAllowedInstance(t *testing.T) {
	allowedMap := map[string]int{
		"grey":   1, // dummy values
		"apple":  2,
		"gopher": 3,
		"vrt":    4,
		"port":   5,
		"sigma":  12,
	}

	allowedFilter := NewAllowedInstance(allowedMap)

	if allowedFilter == nil {
		t.Fatal("Expected instance not to be nil")
	}

	if !reflect.DeepEqual(allowedFilter.AllowedStrings, allowedMap) {
		t.Errorf("NewAllowedInstance returned wrong map: returned %v , expected %v", allowedFilter.AllowedStrings, allowedMap)
	}
}

func TestAllowed_IsValid(t *testing.T) {
	allowedStrings := map[string]int{
		"grey":   1, // dummy values
		"apple":  2,
		"gopher": 3,
		"vrt":    4,
		"port":   5,
		"sigma":  12,
	}

	allowedFilter := NewAllowedInstance(allowedStrings)

	testCases := []struct {
		input    string
		expected bool
	}{
		{"grey", true},
		{"apple", true},
		{"sigma", true},
		{"vrt", true},
		{"ivan", false},
		{"", false},
	}

	for _, tc := range testCases {
		result := allowedFilter.IsValid(tc.input)
		if result != tc.expected {
			t.Errorf("IsValid(%q) = %v : expected %v", tc.input, result, tc.expected)
		}
	}
}
