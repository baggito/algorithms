package main

import "fmt"

func main() {
	arr := []int{0, 1, 2, 3, 4, 5, 6}
	findSmallestMissing(arr, 0, len(arr)-1)
	fmt.Println("The end of algorithm")
}

// Function to find the smallest missing element in a sorted
// array of distinct non-negative integers
func findSmallestMissing(arr []int, left int, right int) int {

	// base condition
	if left > right {
		return left
	}

	mid := left + (right - left)/2

	// if the mid-index matches with its value, then the mismatch
	// lies on the right half
	if arr[mid] == mid {
		return findSmallestMissing(arr, mid + 1, right)
	} else {
		// mismatch lies on the left half
		return findSmallestMissing(arr, left, mid - 1)
	}
}