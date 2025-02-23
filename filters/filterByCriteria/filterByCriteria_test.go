package filterByCriteria

import (
	"filters/validator"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestNewFilterByCriteria(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockValidator := validator.NewMockFailBasedFilter(controller)
	filterCriteria := NewFilterByCriteria(mockValidator)

	if filterCriteria == nil {
		t.Fatal("Expected instance not to be nil")
	}
}

func TestFilterByCriteria_Accepts(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockValidator := validator.NewMockFailBasedFilter(controller)

	testString1 := "valid"
	testString2 := "invalid"
	mockValidator.EXPECT().IsValid(testString1).Return(true).Times(1)
	mockValidator.EXPECT().IsValid(testString2).Return(false).Times(1)

	filterCriteria := NewFilterByCriteria(mockValidator)

	if !filterCriteria.Accepts(testString1) {
		t.Errorf("Expected %v to be accepted by the filter", testString1)
	}
	if filterCriteria.Accepts(testString2) {
		t.Errorf("Expected %v to not be accepted by the filter", testString2)
	}
}
