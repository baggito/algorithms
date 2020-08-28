package main

import (
	"fmt"
)

// initialize minimum element by infinity and the maximum
// element by minus infinity
var minInt = int(^uint(0) >> 1) // MAX
var maxInt = -minInt - 1        // MIN

func main() {
	// given array
	arr := []int{4, 7, 5, 1, 3}
	// array length
	n := len(arr)

	// if the array has odd number of elements, ignore last
	// element and consider it later
	odd := n % 2
	if odd == 1 {
		n--
	}

	// compare elements in pairs i.e. A[i] and A[i+1]
	for i := 0; i < n; i += 2 {
		maximum, minimum := 0, 0

		// find maximum and minimum among A[i], A[i+1]
		if arr[i] > arr[i+1] { // 1st comparison
			minimum = arr[i+1]
			maximum = arr[i]
		} else {
			minimum = arr[i]
			maximum = arr[i+1]
		}

		// update max
		if maximum > maxInt { // 2nd comparison
			maxInt = maximum
		}

		// update min
		if minimum < minInt { // 3rd comparison
			minInt = minimum
		}
	}

	// handle last element if the array has odd number of elements
	if odd == 1 {
		if arr[n] > maxInt {
			maxInt = arr[n]
		}

		if arr[n] < minInt {
			minInt = arr[n]
		}
	}

	fmt.Println("The minimum element in the array is ", minInt)
	fmt.Println("The maximum element in the array is ", maxInt)

	fmt.Println("The end of the algorithm")
}
