package main

func RecursiveSum(arr []int) int {
	idx := len(arr) - 1
	return RecursiveSumRecursive(arr, idx)
}

func RecursiveSumRecursive(arr []int, idx int) int {
	if idx < 0 {
		return 0
	}
	return arr[idx] + RecursiveSumRecursive(arr, idx-1)
}
