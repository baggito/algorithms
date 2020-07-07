package main

import (
	"fmt"
)

func numberOfWays(arr []int, k int) int {
	// map for storing already considered elements
	var occ = map[int]int{}

	// create map with occurrences
	for i := 0; i < len(arr); i++ {

		if _, ok := occ[arr[i]]; ok {
			occ[arr[i]]++
		} else {
			occ[arr[i]] = 1
		}
	}

	// count of pairs
	countAvg, count := 0, 0

	for v, i := range occ {
		if v == k/2 {
			// C(n,2) = n!/2*(n-2!) => (n-1)n/2
			countAvg += (v * (v - 1)) / 2
		} else if _, ok := occ[k-v]; ok {
			count += occ[k-v] * occ[i]
		}
	}

	return countAvg + count/2
}

func main() {
	// Call numberOfWays() with test cases here
	arr := []int{1, 5, 3, 3, 3}
	k := 6

	fmt.Println(numberOfWays(arr, k))
}
