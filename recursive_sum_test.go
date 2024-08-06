package main

import (
	"testing"
)

func TestRecursiveSum(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{1, 2, 3, 4}, 10},      // Typical case
		{[]int{5, 5, 5, 5}, 20},      // Typical case
		{[]int{10, -10, 10, -10}, 0}, // Mixed positive and negative numbers
		{[]int{}, 0},                 // Edge case: empty slice
		{[]int{0, 0, 0, 0}, 0},       // Edge case: all zeros
		{[]int{-1, -2, -3, -4}, -10}, // All negative numbers
		{[]int{100}, 100},            // Single element
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			result := RecursiveSum(tt.nums)
			if result != tt.expected {
				t.Errorf("RecursiveSum(%v) = %d; want %d", tt.nums, result, tt.expected)
			}
		})
	}
}
