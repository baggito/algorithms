package main

import (
	"fmt"
)

func main() {
	// an unsorted array of integers
	arr := []int{8, 7, 1, 5, 3, 2}
	sum := 10

	findPair(arr, sum)

	fmt.Println("The end of algorithm")
}

func findPair(arr []int, sum int) {
	// create an empty map
	var m = map[int]int{}

	// do for each element
	for i := 0; i < len(arr); i++ {
		// check if pair (arr[i], sum-arr[i]) exists

		// if difference is seen before, print the pair
		if val, ok := m[sum-arr[i]]; ok {
			fmt.Printf("Pair found at index: %d and %d \n", val, i)
			return
		}

		// store index of current element in the map
		m[arr[i]] = i
	}

	fmt.Println("Pair not found")
}
