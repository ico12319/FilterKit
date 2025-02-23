package intersectionOfFilters

import (
	"filters/filter"
	"github.com/golang/mock/gomock"
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
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockFilter1 := filter.NewMockFilter(controller)
	mockFilter2 := filter.NewMockFilter(controller)
	mockFilter3 := filter.NewMockFilter(controller)
	mockFilters := []filter.Filter{mockFilter1, mockFilter2, mockFilter3}

	iF := NewIntersectionOfFiltersInstance()
	iF.Add(mockFilter1)
	iF.Add(mockFilter2)
	iF.Add(mockFilter3)

	if !reflect.DeepEqual(iF.Filters, mockFilters) {
		t.Errorf("IntersectionOfFilters_Add failed expected %v returned %v", mockFilters, iF.Filters)
	}
}

func TestIntersectionOfFilters_Accepts(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockFilter1 := filter.NewMockFilter(controller)
	mockFilter2 := filter.NewMockFilter(controller)

	testString1 := "should be accepted"
	testString2 := "should be rejected"
	mockFilter1.EXPECT().Accepts(testString1).Return(true).Times(1)
	mockFilter2.EXPECT().Accepts(testString1).Return(true).Times(1)
	mockFilter1.EXPECT().Accepts(testString2).Return(true).Times(1)
	mockFilter2.EXPECT().Accepts(testString2).Return(false).Times(1)

	iF := NewIntersectionOfFiltersInstance()
	iF.Add(mockFilter1)
	iF.Add(mockFilter2)

	if !iF.Accepts(testString1) {
		t.Errorf("Expected %v to be accepted by the filter", testString1)
	}
	if iF.Accepts(testString2) {
		t.Errorf("Expected %v to not be accepted by the filter", testString2)
	}
}
