package main

import "fmt"

func main() {
	//timeSeries := []int{1,2,3,4,5,6,7,8,9}
	//duration := 1
	// Output = 9
	timeSeries := []int{1,2,3,4,5}
	duration := 5
	// Output = 9


	fmt.Println(findPoisonedDuration(timeSeries, duration))
}

func findPoisonedDuration(timeSeries []int, duration int) int {
	result := duration

	for i := 0; i < len(timeSeries)-1; i++ {
		if timeSeries[i+1] - timeSeries[i] >= duration {
			result += duration
		} else {
			result += timeSeries[i+1] - timeSeries[i]
		}
	}
	return result
}
