package craftMaxLengthFilter

import (
	"testing"
)

func TestCreateMaxLengthFilter(t *testing.T) {
	f := CreateMaxLengthFilter(10) // dummy value just for the purpose of testing

	testString := "short"
	longString := "veryyyy long string"

	if !f.Accepts(testString) {
		t.Errorf("Expected %v to be accepted by the filter.", testString)
	}

	if f.Accepts(longString) {
		t.Errorf("Expected %v to not be accepted by the filter.", testString)
	}
}
