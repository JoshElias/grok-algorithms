package main

func BinarySearch(orderedList []int, item int) (int, bool) {
	upper := len(orderedList) - 1
	lower := 0
	for lower <= upper {
		// When dividing 2 integers, Go will ignore the
		// floating values in the result
		mid := (upper + lower) / 2
		guess := orderedList[mid]
		if guess == item {
			return mid, true
		}
		// If the guess is too high, every value from the index of the guess
		// to the end of the array must be to high as well. Ignore that side
		// the array in future iterations
		if guess > item {
			upper = mid - 1
		} else {
			// It's too low, apply the same logic and ignore the other end
			// of the array
			lower = mid + 1
		}
	}
	return 0, false
}
