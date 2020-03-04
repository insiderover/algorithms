package sorting

import (
	"reflect"
	"testing"
)

type testCaseInsertionSort struct {
	testSlice     []int
	expectedSlice []int
}

var testDataInsertionSort = []testCaseInsertionSort{
	{[]int{}, []int{}},
	{[]int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
	{[]int{5, 4, 3, 2, 1, 0, -1, -2, -3, -4, -5}, []int{-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5}},
}

func TestInsertionSort(t *testing.T) {
	for _, testCaseData := range testDataInsertionSort {
		InsertionSort(testCaseData.testSlice)
		if !reflect.DeepEqual(testCaseData.expectedSlice, testCaseData.testSlice) {
			t.Error(
				"Slice", testCaseData.testSlice,
				"not equal expected result", testCaseData.expectedSlice)
		}
	}
}
