package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{2,7,11,15, -3, -4}
	fmt.Println(threeSum(nums))
}

func threeSum(nums []int) [][]int {
	result := make([][]int, 0)

	// sort an array. O(nlogn)
	sort.Ints(nums)

	// make a hashmap
	hm := make(map[int]int)

	// fill out hashmap. O(n)
	for k, v := range nums {
		hm[v] = k
	}

	start, end := 0, len(nums)-1

	// traverse nums. O(n)
	for start < end {
		d := 0 - (nums[start] + nums[end])

		if index, ok := hm[d]; ok && index != start && index != end  {
			result = append(result, []int{nums[start], nums[end], d})
			start++
			end--
		} else {
			if d > 0 {
				start++
			} else {
				end--
			}
		}
	}
	return result
}
