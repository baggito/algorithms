package main

import (
	"fmt"
)

func main() {
	// given slice
	arr := []int{6, 0, 8, 2, 3, 0, 4, 0, 1}

	// pivot element
	j := 0

	for i := 0; i < len(arr); i++ {
		if arr[i] != 0 {
			swap(arr, i, j)
			j++
		}
	}

	fmt.Println("Result array: ", arr)
	fmt.Println("The end of algorithm")
}

// Helper function that swap two values in array
func swap(arr []int, i, j int) {
	tmp := arr[i]
	arr[i] = arr[j]
	arr[j] = tmp
}
