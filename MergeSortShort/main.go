package main

import (
	"fmt"
)

func main() {
	array := []int{6, 3, 7, 4, 8, 5, 12, 17, 56, 1, 2, 2, 9, 50}

	fmt.Println(mergeSort(array))
}

func mergeSort(array []int) []int {
	if len(array) == 1 {
		return array
	}
	// find mid element
	mid := len(array) / 2

	return merge(mergeSort(array[:mid]), mergeSort(array[mid:]))
}

func merge(left []int, right []int) []int {
	result := make([]int, 0)

	leftIn, rightIn := 0, 0

	for leftIn < len(left) && rightIn < len(right) {
		if left[leftIn] < right[rightIn] {
			result = append(result, left[leftIn])
			leftIn++
		} else {
			result = append(result, right[rightIn])
			rightIn++
		}
	}

	result = append(result, left[leftIn:]...)
	result = append(result, right[rightIn:]...)

	return result
}
