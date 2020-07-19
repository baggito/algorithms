package main

import (
	"fmt"
)

// Function that returns max of two int values
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {
	// given slice
	arr := []int{7, 0, 4, 2, 5, 0, 6, 4, 0, 5}

	// iterator starting from left side
	left := 0

	// iterator starting from right side
	right := len(arr) - 1

	// total water amount
	water := 0

	maxLeft := arr[left]
	maxRight := arr[right]

	for left < right {
		if arr[left] <= arr[right] {
			left++
			maxLeft = max(maxLeft, arr[left])
			water += maxLeft - arr[left]
		} else {
			right--
			maxRight = max(maxRight, arr[right])
			water += maxRight - arr[right]
		}
	}

	fmt.Println("Total amount of water is: ", water)
	fmt.Println("The end of algorithm")
}
