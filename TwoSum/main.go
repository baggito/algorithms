package main

import "fmt"

func main() {
	nums := []int{2,7,11,15}
	target := 9
	fmt.Println(twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {
	// make a hashmap
	hm := make(map[int]int)
	var dif int

	for k, v := range nums {
		dif = target - v
		if val, ok := hm[dif]; ok {
			return []int{val, k}
		}
		hm[v] = k
	}

	return []int{}
}
