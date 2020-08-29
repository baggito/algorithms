package main

import (
	"fmt"
)

func main() {

	// given array
	arr := []int{5, 4, 5, 5, 3, 1, 2, 2, 4}

	// create an empty map to store frequency of each element
	var hMap = map[int]int{}
	// array of hMap's keys with original order
	var hMapOrder []int

	// traverse the input array and update frequency of each element
	for _, v := range arr {
		if _, ok := hMap[v]; !ok {
			hMap[v] = 1
			// storing elements in array in an initial order (without duplicates)
			hMapOrder = append(hMapOrder, v)
		} else {
			hMap[v] = hMap[v] + 1
		}
	}

	var result []int
	// traverse the distinct array of keys and print each one times equal to hMap[v]
	for _, v := range hMapOrder {
		for i := 0; i < hMap[v]; i++ {
			result = append(result, v)
		}
	}

	fmt.Println(result)
	fmt.Println("The end of the algorithm")
}
