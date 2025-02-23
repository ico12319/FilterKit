package craftIntersectionFilter

import "testing"

func TestCreateIntersectionFilter(t *testing.T) {
	iF := CreateIntersectionFilter("allowed2.txt", "blocked2.txt", 6)
	/* Here is what is contained in allowed.txt :
	gopher
	secret
	helloworld
	admin
	*/
	/* Here is what is contained in allowed.txt :
	admin
	test
	*/
	testString1 := "gopher"     // it is contained in the allowed.txt and is not contained in blocked.txt also the length is 6 so we expect true
	testString2 := "admin"      // contained in allowed.txt but also present in blocked.txt so we expect false
	testString3 := "helloworld" // contained in allowed.txt but its length exceeds 6 so we expect false

	if !iF.Accepts(testString1) {
		t.Errorf("Expected %v to be accepted by the filter", testString1)
	} else if iF.Accepts(testString2) {
		t.Errorf("Expected %v to not be accepted by the filter", testString2)
	} else if iF.Accepts(testString3) {
		t.Errorf("Expected %v to not be accepted by the filter", testString3)
	}
}
