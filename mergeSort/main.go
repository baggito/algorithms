package main

import (
	"fmt"
)

func main() {
	// given unsorted array
	arr := []int{1, 5, 2, 7, 4, 8, 3, 9, 19, 10, -2, -1, 0}
	fmt.Println("Unsorted array: ", arr)

	sortedArr := mergeSort(arr)
	fmt.Println("Sorted array:   ", sortedArr)

	fmt.Println("The end of algorithm")
}

func mergeSort(slice []int) []int {
	// we have reached the point when all slices have length 1
	if len(slice) < 2 {
		return slice
	}

	// find current slice mid point
	mid := len(slice) / 2

	return merge(mergeSort(slice[:mid]), mergeSort(slice[mid:]))
}

func merge(left, right []int) []int {
	// the size of new slice that contains merge result of left and right slices
	size := len(left) + len(right)
	// new slice for merge
	slice := make([]int, size, size)
	// iterators
	l, r := 0, 0

	// go through merge slice
	for s := 0; s < size; s++ {
		if l > len(left)-1 && r <= len(right)-1 {
			slice[s] = right[r]
			r++
		} else if l <= len(left)-1 && r > len(right)-1 {
			slice[s] = left[l]
			l++
		} else if left[l] <= right[r] {
			slice[s] = left[l]
			l++
		} else {
			slice[s] = right[r]
			r++
		}
	}
	return slice
}
