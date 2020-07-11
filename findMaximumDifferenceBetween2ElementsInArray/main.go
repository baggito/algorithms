package main

import (
	"fmt"
)

func main() {
	// given array
	arr := []int{2, 7, 9, 5, 1, 3, 5}

	max := arr[0]
	min := arr[0]

	maxIndex := 0

	for i, v := range arr {
		if max < v {
			max = v
			maxIndex = i
		}

		if (min > v) && (i < maxIndex) {
			min = v
		}
	}

	fmt.Printf("The pair is (%d, %d) \n", min, max)
	fmt.Println("The end of algorithm")
}
