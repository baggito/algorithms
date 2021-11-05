package main

import (
	"fmt"
	"math"
)

type Point struct {
	X float64
	Y float64
}

type Vertex struct {
	X float64
	Y float64
	D float64
}

func main() {

	points := []Point{{9,2}, {4,5}, {3,6}, {0,0}, {5,1}}
	point  := Point{2,2}
	//k := 2

	vertexes := make([]Vertex, 0)

	// calculate vertexes
	for _, v := range points {
		// calculate a distance
		distance := math.Round(math.Sqrt(math.Pow((v.X - point.X), 2) + math.Pow((v.Y - point.Y), 2)))
		vertexes = append(vertexes, Vertex{v.X, v.Y, distance})
	}

	fmt.Printf("Vertexes: %+v \n", mergeSort(vertexes))
	fmt.Println("The End")
}

func mergeSort(slice []Vertex) []Vertex {
	// we have reached the point when all slices have length 1
	if len(slice) < 2 {
		return slice
	}

	// find current slice mid point
	mid := len(slice) / 2

	return merge(mergeSort(slice[:mid]), mergeSort(slice[mid:]))
}

func merge(left, right []Vertex) []Vertex {
	// the size of new slice that contains merge result of left and right slices
	size := len(left) + len(right)
	// new slice for merge
	slice := make([]Vertex, size, size)
	// iterators
	l, r := 0, 0

	// go through merge slice
	for s := 0; s < size; s++ {
		if l > len(left)-1 && r <= len(right)-1 {
			slice[s] = right[r]
			r++
		} else if l <= len(left)-1 && r > len(right)-1 {
			slice[s] = left[l]
			l++
		} else if left[l].D <= right[r].D {
			slice[s] = left[l]
			l++
		} else {
			slice[s] = right[r]
			r++
		}
	}
	return slice
}