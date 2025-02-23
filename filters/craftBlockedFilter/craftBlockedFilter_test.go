package craftBlockedFilter

import "testing"

func TestCreateBlockedFilter(t *testing.T) {
	bF := CreateBlockedFilter("blocked.txt")
	/* Here is what is contained in allowed.txt :
	admin -> expected return value of the object created -> false
	test -> exptected return value of the object created -> false
	*/
	testString1 := "admin"
	testString2 := "test"
	testString3 := "this should return true"
	if bF.Accepts(testString1) {
		t.Errorf("Expected %v to not be accepted by the filter", testString1)
	} else if bF.Accepts(testString2) {
		t.Errorf("Expected %v to not be accepted by the filter", testString2)
	} else if !bF.Accepts(testString3) {
		t.Errorf("Expected %v to be accepted by the filter", testString3)
	}
}
