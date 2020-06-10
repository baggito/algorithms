package main

import (
	"fmt"
)

func main() {
	// given a sorted array of distinct	non-negative integers
	arr := []int{0, 1, 2, 6, 9, 11, 15}
	//arr := []int{1, 2, 3, 4, 6, 9, 11, 15}
	//arr := []int{0, 1, 2, 3, 4, 5, 6}

	// find the smallest missing element
	findElement(arr)

	fmt.Println("The end of algorithm")
}

func findElement(arr []int) {
	// first smallest element should be 0
	if arr[0] != 0 {
		fmt.Println("The smallest missing element is: ", 0)
		return
	}

	var i int
	for i = 0; i < len(arr)-1; i++ {
		// this condition checks if there is any number between
		if arr[i+1]-arr[i] > 1 {
			fmt.Println("The smallest missing element is: ", arr[i]+1)
			return
		}
	}

	// if there is no any missing element then missing one will be the last number + 1
	fmt.Println("The smallest missing element is: ", arr[i]+1)
	return
}
