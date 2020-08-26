package main

import (
	"fmt"
)

func splitByPoints(ranges [][]int, points []int) [][]int {
	var result []int
	// slice of start points of all ranges [0, 1] [4, 8], [9, 13]: it will be [0, 4, 9]
	var startPoints = map[int]int{}
	// slice of end points of all ranges [0, 1] [4, 8], [9, 13]: it will be [1, 8, 13]
	var endPoints = map[int]int{}

	// generate a new one-dimensional array
	var rangeArr []int
	for _, v := range ranges {
		rangeArr = append(rangeArr, v[0])
		rangeArr = append(rangeArr, v[1])
		// initialize startPoints and endPoints with corresponding values
		startPoints[v[0]] = v[0]
		endPoints[v[1]] = v[1]
	}

	var i, j int
	for i, j = 0, 0; i < len(rangeArr) && j < len(points); {
		if rangeArr[i] < points[j] {
			result = append(result, rangeArr[i])
			i++
		} else if rangeArr[i] == points[j] { // avoiding duplicates
			result = append(result, rangeArr[i])
			i++
			j++
		} else {
			result = append(result, points[j])
			j++
		}
	}

	for i < len(rangeArr) {
		result = append(result, rangeArr[i])
		i++
	}

	var finalResult [][]int

	for i := 0; i < len(result)-1; i++ {
		_, okE := endPoints[result[i]]
		_, okS := startPoints[result[i+1]]

		// if result[i] is an end point  and result[i+1] is a start point then we should not
		// include them as a range pair
		if !(okE && okS) {
			finalResult = append(finalResult, []int{result[i], result[i+1]})
		}
	}

	return finalResult
}

func main() {

	// slice of ranges
	ranges := [][]int{
		{0, 1},
		{4, 8},
		{9, 13},
	}

	// slice of points
	points := []int{1, 5, 6, 11, 15}

	result := splitByPoints(ranges, points)

	fmt.Println(result)
	fmt.Println("The end of the algorithm")

}
