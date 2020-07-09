package main

import (
	"fmt"
)

func main() {
	// given slice
	arr := []int{2, 8, 7, 2, 2, 5, 2, 3, 1, 2, 2} // majority element is 2
	//arr := []int{1,2,3,1} // there is no majority element

	var m, j int
	for i, v := range arr {
		if i == 0 {
			m = v
			j++
		} else if v == m {
			m = v
			j++
		} else {
			j--
		}
	}

	// majority element occurrence
	count := 0
	// validate majority element
	for _, v := range arr {
		if v == m {
			count++
		}
	}

	if count > len(arr)/2 {
		fmt.Println("The majority is: ", m)
	} else {
		fmt.Println("There is no majority element")
	}
	fmt.Println("The end of the algorithm")
}
