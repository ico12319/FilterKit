package unionOfFilters

import (
	"filters/filter"
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"
)

func TestNewUnionOfFiltersInstance(t *testing.T) {
	instance := NewUnionOfFiltersInstance()

	if instance == nil {
		t.Fatal("Expected instance not to be nil")
	}
}

func TestUnionOfFilters_Add(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockFilter1 := filter.NewMockFilter(controller)
	mockFilter2 := filter.NewMockFilter(controller)
	mockFilter3 := filter.NewMockFilter(controller)
	mockFilters := []filter.Filter{mockFilter1, mockFilter2, mockFilter3}

	uF := NewUnionOfFiltersInstance()
	uF.Add(mockFilter1)
	uF.Add(mockFilter2)
	uF.Add(mockFilter3)

	if !reflect.DeepEqual(uF.Filters, mockFilters) {
		t.Errorf("UnionOfFilters_Add failed expected %v returned %v", mockFilters, uF.Filters)
	}

}

func TestUnionOfFilters_Accepts(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockFilter1 := filter.NewMockFilter(controller)
	mockFilter2 := filter.NewMockFilter(controller)

	testString1 := "apple"
	testString2 := "gopher"
	mockFilter1.EXPECT().Accepts(testString1).Return(false).Times(1)
	mockFilter2.EXPECT().Accepts(testString1).Return(false).Times(1)

	mockFilter1.EXPECT().Accepts(testString2).Return(true).Times(1)
	mockFilter2.EXPECT().Accepts(testString2).Return(false).Times(0)

	uF := NewUnionOfFiltersInstance()
	uF.Add(mockFilter1)
	uF.Add(mockFilter2)

	if uF.Accepts(testString1) {
		t.Errorf("Expected %v to not be accepted by the filter", testString1)
	}

	if !uF.Accepts(testString2) {
		t.Errorf("Expected %v to be accepted by the filter", testString2)
	}

}
