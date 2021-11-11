package main

import (
	"container/heap"
	"fmt"
	"sort"
)

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	nums := []int{3, 1, 7, 5}

	sort.Slice(nums, func(i, j int) bool { // O (n log n)
		return nums[i] < nums[j]
	})

	fmt.Println(nums)

	// initialize the heap data structure
	h := &IntHeap{}

	heap.Push(h, 3)
	heap.Push(h, 1)
	heap.Push(h, 7)
	heap.Push(h, 5)

	// add all the values to heap, O(n log n)
	//for _, val := range nums { // O(n)
	//	heap.Push(h, val) // O(log n)
	//}

	// print all the values from the heap
	// which should be in ascending order
	for i := 0; i < len(nums); i++ {
		fmt.Printf("%d,", heap.Pop(h).(int))
	}
}
