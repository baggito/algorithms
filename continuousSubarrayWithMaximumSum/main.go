package main

import (
	"fmt"
)

func main() {
	// Given array
	arr := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}

	// stores maximum sum subarray found so far
	maxSoFar := arr[0]

	// stores maximum sum of subarray ending at current position
	maxEndingHere := arr[0]

	// stores end-points of maximum sum subarray found so far
	start, end := 0, 0

	// stores starting index of a positive sum sequence
	beg := 0

	for i, v := range arr {
		// update maximum sum of subarray "ending" at index i
		maxEndingHere += v

		// if maximum sum is negative, set it to 0
		if maxEndingHere < v {
			maxEndingHere = v
			beg = i
		}

		// update result if current subarray sum is found to be greater
		if maxSoFar < maxEndingHere {
			maxSoFar = maxEndingHere
			start = beg
			end = i
		}
	}

	fmt.Println("The sum of contiguous subarray with the largest sum is ", maxSoFar)
	fmt.Println("The contiguous subarray with the largest sum is ", arr[start:end+1])
	fmt.Println("The end of algorithm")
}
