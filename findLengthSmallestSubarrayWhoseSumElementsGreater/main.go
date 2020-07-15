package main

import (
	"fmt"
	"math"
)

func main() {
	// given positive array
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	k := 20

	// stores the current window sum
	windowSum := 0

	// stores the result
	subarrayLen := math.MaxFloat64

	// stores window's starting index
	left := 0

	// maintain a sliding window [left..right]
	for right := range arr {
		// include current element in the window
		windowSum += arr[right]

		for windowSum > k && left <= right {
			// update the result if current window's length is less
			// than minimum found so far
			subarrayLen = math.Min(subarrayLen, float64(right-left+1))

			// remove elements from the window's left side till window
			// becomes stable again
			windowSum -= arr[left]
			left++
		}
	}

	if subarrayLen != math.MaxFloat64 {
		fmt.Printf("Smallest sub-array length is %d \n", int(subarrayLen))
	} else {
		fmt.Printf("No sub-array exists \n")
	}

	fmt.Println("The end of algorithm")
}
