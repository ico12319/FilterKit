package craftAllowedFilter

import "testing"

func TestCreateAllowedFilter(t *testing.T) {
	aF := CreateAllowedFilter("allowed.txt")
	/* Here is what is contained in allowed.txt :
	gopher -> expected return value of the object created -> true
	secret -> expected return value of the object created -> true
	helloworld -> expected return value of the object created -> true
	admin -> expected return value of the object created -> true
	*/

	testString1 := "gopher"
	testString2 := "secret"
	testString3 := "this should return false"
	if !aF.Accepts(testString1) {
		t.Errorf("Expected %v to be accepted by the filter", testString1)
	} else if !aF.Accepts(testString2) {
		t.Errorf("Expected %v to be accepted by the filter", testString2)
	} else if aF.Accepts(testString3) {
		t.Errorf("Expected %v to not be accepted by the filter", testString3)
	}
}
