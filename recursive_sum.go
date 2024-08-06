package main

func RecursiveSum(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	return arr[0] + RecursiveSum(arr[1:])
}
