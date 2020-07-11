package main

import (
	"fmt"
)

func main() {
	// Given array
	arr := []int{3, 5, 8, 4, 5, 9, 10, 8, 5, 3, 4}

	endIndex, maxLen := 0, 0

	for i := 0; i+1 < len(arr); {
		// check for Longest Bitonic Subarray starting at A[i]

		// reset length to 1
		tmpLen := 1

		// run till sequence is increasing
		for i+1 < len(arr) && arr[i] < arr[i+1] {
			i++
			tmpLen++
		}

		// run till sequence is decreasing
		for i+1 < len(arr) && arr[i] > arr[i+1] {
			i++
			tmpLen++
		}

		// update Longest Bitonic Subarray if required
		if maxLen < tmpLen {
			maxLen = tmpLen
			endIndex = i
		}
	}

	fmt.Println("The longest bitonic sub-array is: ", arr[endIndex-maxLen+1:endIndex+1])
	fmt.Println("The end of algorithm")
}
