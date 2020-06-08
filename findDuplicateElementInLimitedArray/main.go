package main

import (
	"fmt"
)

func main() {
	// given limited range array with size n, where array contains elements between 1 to n-1
	arr := []int{1, 2, 3, 4, 4}

	findDuplicate(&arr)

	fmt.Println("The end of algorithm")
}

// find sum of all element and find difference between it and all elements which are supposed to be there.
func findDuplicate(arr *[]int) {
	// get a length of given array
	n := len(*arr)
	// find the sum of elements of 1 2... n
	sum1 := n * (n + 1) / 2

	// find the sum of elements of given array
	sum2 := 0
	for i := 0; i < n; i++ {
		sum2 += (*arr)[i]
	}

	// this give us exact duplicate value
	// lets say we have {1 2 3 4 5}, so the sum is 15
	// but {1 2 3 4 4} has a duplicate so the sum should be less than 15, so difference between 15 and 14 will
	// show the offset 15 - 14 = 1, so duplicate less than n (array length) by offset. 4 = 5 - (15-14)
	duplicate := n - (sum1 - sum2)

	fmt.Println("The duplicate is: ", duplicate)
}
