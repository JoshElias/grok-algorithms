package main

func BinarySearch(orderedList []int, item int) (int, bool) {
	upper := len(orderedList) - 1
	lower := 0
	for lower <= upper {
		mid := (upper + lower) / 2
		guess := orderedList[mid]
		if guess == item {
			return mid, true
		}
		if guess > item {
			upper = mid - 1
		} else {
			lower = mid + 1
		}
	}
	return 0, false
}
