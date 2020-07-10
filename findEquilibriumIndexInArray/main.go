package main

import (
	"fmt"
)

func main() {
	// given array
	arr := []int{0, -3, 5, -4, -2, 3, 1, 0}

	// get sum of array
	arrSum := sum(arr)

	leftSum := 0

	// slice of Equilibrium indexes
	var eq []int

	for i := 1; i < len(arr)-1; i++ {
		leftSum += arr[i-1]

		if leftSum == arrSum-arr[i]-leftSum {
			eq = append(eq, i)
		}
	}

	if len(eq) != 0 {
		fmt.Println("Equilibrium indexes: ", eq)
	} else {
		fmt.Println("There is no Equilibrium indexes")
	}

	fmt.Println("The end of algorithm")
}

func sum(arr []int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}

	return sum
}
