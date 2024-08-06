package main

import (
	"testing"
)

func TestGCD(t *testing.T) {
	tests := []struct {
		a, b     int
		expected int
	}{
		{1680, 640, 80}, // Typical case
		{54, 24, 6},     // Typical case
		{48, 18, 6},     // Typical case
		{0, 5, 5},       // Edge case: one number is zero
		{5, 0, 5},       // Edge case: one number is zero
		{0, 0, 0},       // Edge case: both numbers are zero
		{7, 7, 7},       // Identical numbers
		{13, 17, 1},     // Prime numbers
		{-54, 24, 6},    // Negative number
		{54, -24, 6},    // Negative number
		{-54, -24, 6},   // Both numbers negative
	}

	for _, test := range tests {
		result := GCD(test.a, test.b)
		if result != test.expected {
			t.Errorf("GCD(%d, %d) = %d; want %d", test.a, test.b, result, test.expected)
		}
	}
}
