package main

import (
	"fmt"
)

func main() {
	// given array of integers
	arr := []int{3, 4, -7, 3, 1, 3, 1, -4, -2, -2}

	// Complexity O(n^3)
	//findSubarrayO3(arr)

	// Complexity O(n^2)
	findSubarrayO2(arr)

	fmt.Println("The end of algorithm")
}

// Complexity O(n^3)
func findSubarrayO3(arr []int) {
	var res []int

	// traverse the given array
	for i := 0; i < len(arr); i++ {
		sum := 0

		// traverse the given array for sub arrays
		for j := i; j < len(arr); j++ {
			// sum of elements so far
			sum += arr[j]
			res = append(res, arr[j])

			if sum == 0 {
				fmt.Println(res)
			}
		}
		res = []int{}
	}
}

// Utility function to insert <key, value> into the Multimap
func insert(hMap map[int][]int, k int, v int) {
	// if the key is seen for the first time, initialize the list
	if _, ok := hMap[k]; !ok {
		hMap[k] = []int{}
	}

	hMap[k] = append(hMap[k], v)
}

// Complexity O(n^2)
func findSubarrayO2(arr []int) {
	// create an empty Multi-map to store ending index of all
	// sub-arrays having same sum
	var hMap = map[int][]int{}

	// insert (0, -1) pair into the map to handle the case when
	// sub-array with 0 sum starts from index 0
	insert(hMap, 0, -1)

	sum := 0

	// traverse the given array
	for i := 0; i < len(arr); i++ {
		// sum of elements so far
		sum += arr[i]

		// if sum is seen before, there exists at-least one
		// sub-array with 0 sum
		if list, ok := hMap[sum]; ok {
			// find all sub-arrays with same sum
			for _, v := range list {
				fmt.Printf("Subarray [%d .. %d] \n", v+1, i)
			}
		}

		// insert (sum so far, current index) pair into the Multi-map
		insert(hMap, sum, i)
	}
}
