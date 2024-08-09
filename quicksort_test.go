package main

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{input: []int{}, expected: []int{}},
		{input: []int{1}, expected: []int{1}},
		{input: []int{2, 1}, expected: []int{1, 2}},
		{input: []int{5, 3, 8, 4, 2}, expected: []int{2, 3, 4, 5, 8}},
		{input: []int{10, 7, 8, 9, 1, 5}, expected: []int{1, 5, 7, 8, 9, 10}},
		{input: []int{3, -1, 0, -3, 2, 1}, expected: []int{-3, -1, 0, 1, 2, 3}},
		{input: []int{5, 5, 5, 5, 5}, expected: []int{5, 5, 5, 5, 5}},
		{input: []int{9, 7, 5, 3, 1}, expected: []int{1, 3, 5, 7, 9}},
	}

	for _, test := range tests {
		sorted := QuickSort(test.input)
		if !reflect.DeepEqual(sorted, test.expected) {
			t.Errorf("QuickSort(%v) = %v; want %v", sorted, test.input, test.expected)
		}
	}
}
