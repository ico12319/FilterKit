package lengthBlocker

import (
	"testing"
)

func TestNewLengthBlocker(t *testing.T) {
	dummyLength := 21
	lBlockerInstance := NewLengthBlocker(dummyLength) // dummy number, for the purpose of testing
	if lBlockerInstance == nil {
		t.Fatal("Expected instance not to be nil")
	}
	if lBlockerInstance.MaxLength != dummyLength {
		t.Errorf("Incorrect maxLength : expected %v, returned %v", dummyLength, lBlockerInstance.MaxLength)
	}
}

func TestLengthBlocker_IsValid(t *testing.T) {
	dummyMaxLength := 3

	lBlockerInstance := NewLengthBlocker(dummyMaxLength)
	testCases := []struct {
		input    string
		expected bool
	}{
		{"hello", false},
		{"iv", true},
		{"er", true},
		{"k", true},
		{"", true},
		{"dog", true},
		{"akjflahgla", false},
		{"falhhaluaghaopfhagoayrqogfakg", false},
	}

	for _, tc := range testCases {
		result := lBlockerInstance.IsValid(tc.input)
		if result != tc.expected {
			t.Errorf("IsValid(%q) = %v, expected %v", tc.input, result, tc.expected)
		}
	}
}
