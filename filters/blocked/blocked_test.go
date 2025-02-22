package blocked

import (
	"reflect"
	"testing"
)

func TestNewBlockedInstance(t *testing.T) {
	blockedStrings := map[string]int{
		"apple": 1, // dummy values
		"black": 2,
		"kusha": 3,
		"chica": 13,
		"kusa":  14,
		"admin": 4,
		"love":  6,
	}
	instance := NewBlockedInstance(blockedStrings)

	if instance == nil {
		t.Fatal("Expected instance not to be nil")
	}
	if !reflect.DeepEqual(instance.BlockedStrings, blockedStrings) {
		t.Errorf("NewBlockedInstance returned wrong map: returned %v, expected %v", instance.BlockedStrings, blockedStrings)
	}
}

func TestBlocked_IsValid(t *testing.T) {
	blockedStrings := map[string]int{
		"apple": 1, // dummy values
		"black": 2,
		"kusha": 3,
		"chica": 13,
		"kusa":  14,
		"admin": 4,
		"love":  6,
	}

	blockedFilter := NewBlockedInstance(blockedStrings)

	testCases := []struct {
		input    string
		expected bool
	}{
		{"apple", false},
		{"black", false},
		{"kusha", false},
		{"chica", false},
		{"admin", false},
		{"love", false},
		{"", true},
		{"gopher", true},
	}

	for _, tc := range testCases {
		result := blockedFilter.IsValid(tc.input)
		if result != tc.expected {
			t.Errorf("IsValid(%q) = %v expected : %v", tc.input, result, tc.expected)
		}
	}

}
