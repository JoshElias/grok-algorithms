package main

import (
	"reflect"
	"testing"
)

func TestSelectiveSort(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{input: []int{5, 3, 6, 2, 10}, expected: []int{2, 3, 5, 6, 10}},
		{input: []int{1, 2, 3, 4, 5}, expected: []int{1, 2, 3, 4, 5}},
		{input: []int{5, 4, 3, 2, 1}, expected: []int{1, 2, 3, 4, 5}},
		{input: []int{3, 3, 3, 3, 3}, expected: []int{3, 3, 3, 3, 3}},
		{input: []int{}, expected: []int{}},
		{input: []int{1}, expected: []int{1}},
	}

	for _, test := range tests {
		result := SelectiveSort(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("SelectiveSort(%v) = %v; want %v", test.input, result, test.expected)
		}
	}
}
