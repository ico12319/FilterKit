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
	allowedMap := map[string]int{
		"grey":   1, // dummy values
		"apple":  2,
		"gopher": 3,
		"vrt":    4,
		"port":   5,
		"sigma":  12,
	}
	allowedFilter := NewAllowedInstance(allowedMap)

	testString1 := "grey"   // should return true
	testString2 := "apple"  // should return true
	testString3 := "hfagqq" // should return false

	if !allowedFilter.IsValid(testString1) {
		t.Errorf("Expected %v to be valid", testString1)
	}
	if !allowedFilter.IsValid(testString2) {
		t.Errorf("Expected %v to be valid", testString2)
	}
	if allowedFilter.IsValid(testString3) {
		t.Errorf("Expected %v to not be valid", testString3)
	}

}
