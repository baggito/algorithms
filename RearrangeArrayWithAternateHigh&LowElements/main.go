package main

import (
	"fmt"
)

func main() {
	// given array
	//arr := []int{6,9,2,5,7,4,1,3}
	arr := []int{1, 2, 3, 4, 5, 6, 7}

	rearrange(arr)

	fmt.Println("The end of algorithm")
}

func rearrange(arr []int) {
	// go thorough given slice
	for i := 1; i < len(arr); i += 2 {
		// if previous element is greater then the current element, we swap the elements
		if arr[i-1] > arr[i] {
			swap(arr, i-1, i)
		}

		// if next element is greater then the current element, we swap the elements
		if i+1 < len(arr) && arr[i+1] > arr[i] {
			swap(arr, i+1, i)
		}
	}

	fmt.Println(arr)
}

// Method to swap to elements in slice
func swap(arr []int, i, j int) {
	tmp := arr[i]
	arr[i] = arr[j]
	arr[j] = tmp
}
