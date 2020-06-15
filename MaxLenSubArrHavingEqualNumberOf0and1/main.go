package main

import (
	"fmt"
)

func main() {
	// given binary array
	arr := []int{0, 0, 1, 0, 1, 0, 0}

	// find sub arrays
	findSubArr(arr)

	fmt.Println("The end of algorithm")
}

func findSubArr(arr []int) {
	// map to store ending index of first sub-array
	m := map[int]int{}

	// insert (0, -1) pair into the set to handle the case when
	// sub-array with sum 0 starts from index 0
	m[0] = -1

	// len stores the maximum length of sub-array with sum 0
	length := 0

	// stores ending index of maximum length sub-array having sum 0
	endingIndex := -1

	sum := 0

	for i := 0; i < len(arr); i++ {
		// sum of elements so far (replace 0 with -1)
		switch arr[i] {
		case 0:
			sum += -1
		default:
			sum += 1
		}

		// if sum is seen before
		if sumIndex, ok := m[sum]; ok {

			// update length and ending index of maximum length
			// sub-array having sum 0
			if length < i-sumIndex {
				length = i - sumIndex
				endingIndex = i
			}
		} else {
			// if sum is seen for first time, insert sum with its
			// index into the map
			m[sum] = i
		}
	}

	if endingIndex != -1 {
		fmt.Printf("[%d, %d] \n", endingIndex-length+1, endingIndex)
	} else {
		fmt.Println("There is no sub-array")
	}
}
