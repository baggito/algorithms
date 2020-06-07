package main

import (
	"fmt"
)

func main() {
	// given binary array
	arr := []int{1, 0, 1, 0, 1, 0, 0, 1}

	// sort
	sort(arr)

	fmt.Println("The end of algorithm")
}

// sort array with O(n) complexity
func sort(arr []int) {
	start := 0
	end := len(arr) - 1

	for start < end {
		if arr[start] == 1 {
			end--
		}
		if arr[start] == 0 {
			start++
		}
		swap(&arr, start, end)
		start++
		end--
	}

	fmt.Println(arr)
}

// utility function to swap two elements of given array
func swap(arr *[]int, i, j int) {
	tmp := (*arr)[i]
	(*arr)[i] = (*arr)[j]
	(*arr)[j] = tmp
}
