package binary_search

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
		{[]int{1, 2, 3, 4, 5}, 3, 2, true},
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
