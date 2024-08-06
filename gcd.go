package main

// This uses the Euclidean Algorithm to get the greatest common divisor
// It uses Divide and Conquer to continually reduce the problem

func GCD(a int, b int) int {
	// These are our 2 base cases
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}

	// Now we deal with the recursive case

	// Get the remainder when dividing with the min
	// Also, use the absolute value of a and b so the GCD is also absolute
	min, max := GetMinAndMax(Abs(a), Abs(b))
	remainder := max % min
	return GCD(min, remainder)
}

func GetMinAndMax(a int, b int) (int, int) {
	if a < b {
		return a, b
	}
	return b, a
}

func Abs(num int) int {
	if num > 0 {
		return num
	}
	return num * -1
}
