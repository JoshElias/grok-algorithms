package main

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		orderedList    []int
		item           int
		expectedIdx    int
		expectedExists bool
	}{
		// Basic cases
		{[]int{1, 2, 3, 4, 5}, 3, 2, true},
		{[]int{1, 2, 3, 4, 5}, 1, 0, true},
		{[]int{1, 2, 3, 4, 5}, 5, 4, true},
		{[]int{1, 2, 3, 4, 5}, 6, 0, false},
		{[]int{}, 1, 0, false},

		// Larger list
		{[]int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23, 25, 27, 29, 31, 33, 35, 37, 39}, 25, 12, true},
		{[]int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23, 25, 27, 29, 31, 33, 35, 37, 39}, 2, 0, false},

		// Edge cases
		{[]int{1}, 1, 0, true},
		{[]int{1}, 2, 0, false},
		{[]int{1, 2}, 1, 0, true},
		{[]int{1, 2}, 2, 1, true},
		{[]int{1, 2}, 3, 0, false},

		// Negative numbers
		{[]int{-10, -5, 0, 5, 10}, -5, 1, true},
		{[]int{-10, -5, 0, 5, 10}, 0, 2, true},
		{[]int{-10, -5, 0, 5, 10}, 10, 4, true},
		{[]int{-10, -5, 0, 5, 10}, -6, 0, false},
	}

	for _, test := range tests {
		idx, exists := BinarySearch(test.orderedList, test.item)
		if idx != test.expectedIdx || exists != test.expectedExists {
			t.Errorf("BinarySearch(%v, %d) = (%d, %v); want (%d, %v)",
				test.orderedList,
				test.item,
				idx,
				exists,
				test.expectedIdx,
				test.expectedExists,
			)
		}
	}
}
