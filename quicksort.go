package main

import (
	"math/rand"
)

func QuickSort(arr []int) []int {
	arrLen := len(arr)

	// Already sorted
	if arrLen < 2 {
		return arr
	}

	// Swap if needed
	if arrLen == 2 {
		// Make copy since we need to mutate
		arrCopy := make([]int, arrLen)
		copy(arrCopy, arr)
		if arrCopy[0] > arrCopy[1] {
			arrCopy[0], arrCopy[1] = arrCopy[1], arrCopy[0]
		}
		return arrCopy
	}

	// Pivot and Split it...Well, partition it...
	random := rand.Intn(arrLen)
	pivot := arr[random]

	var left []int
	var right []int

	// Put the smaller and equal values on the left
	// And the larger ones on the right
	for _, v := range arr {
		if v <= pivot {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}

	// Sort both sides
	left = QuickSort(left)
	right = QuickSort(right)

	// Combine back into a single array
	sorted := make([]int, 0)
	sorted = append(sorted, left...)
	sorted = append(sorted, pivot)
	sorted = append(left, right...)
	return sorted
}
