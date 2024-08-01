package main

import (
	"math"
)

func SelectiveSort(list []int) []int {
	newArr := make([]int, len(list))
	arrCopy := make([]int, len(list))
	copy(arrCopy, list)
	for i := range arrCopy {
		smallestIdx := findSmallestIdx(arrCopy)
		newArr[i] = arrCopy[smallestIdx]
		newCopy, _ := removeUnordered(arrCopy, smallestIdx)
		arrCopy = newCopy

	}
	return newArr
}

func removeUnordered(arr []int, idx int) ([]int, bool) {
	if idx < 0 || idx > len(arr) {
		return nil, false
	}
	arr[idx] = arr[len(arr)-1]
	return arr[:len(arr)-1], true
}

func findSmallestIdx(list []int) int {
	smallest := math.MaxInt64
	smallest_idx := 0
	for i, num := range list {
		if num < smallest {
			smallest = num
			smallest_idx = i
		}
	}
	return smallest_idx
}
