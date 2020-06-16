package main

import (
	"fmt"
)

func main() {
	// given array of 0,1,2
	arr := []int{0, 1, 2, 2, 1, 0, 0, 2, 0, 1, 1, 0}

	// sort an array
	sortArr(arr)

	fmt.Println("The end of algorithm")
}

func sortArr(arr []int) {
	// pivot value
	pivot := 1
	start := 0
	mid := 0
	end := len(arr) - 1

	for mid <= end {
		if arr[mid] < pivot {
			// if mid element is less than pivot then it should go to the front part of array
			swap(arr, start, mid)
			start++
			mid++
		} else if arr[mid] == pivot {
			// if mid element equal to pivot then we skip it
			mid++
		} else {
			// if mid element greater than pivot then it should go to the end part of array
			swap(arr, end, mid)
			end--
		}
	}

	fmt.Println(arr)
}

func swap(arr []int, i, j int) {
	tmp := arr[i]
	arr[i] = arr[j]
	arr[j] = tmp
}
