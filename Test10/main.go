package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	// our main array for asteroids positions
	var arr [][]int

	arr = initialize(arr)
	fmt.Printf("Array: %v \n", arr)

	// Going through each row  0 => [-1 0 -1 -1 0]...
	// y = 0, row = [-1 0 -1 -1 0]
	for y, row := range arr {
		// x = 0, v = -1
		for x, val := range row {
			// we found asteroid, current one
			if val == 0 {
				//fmt.Printf("Parent: (%d, %d) \n", x, y)
				// map to store k and b (y = kx + b), where k - key and b will be stored in slice(float64)
				hmap := map[float64][]float64{}

				// count - total count of available asteroids
				// countX maximum might be 1, available asteroids on the same X axis
				// countY maximum might be 1, available asteroids on the same Y axis
				count, countXP, countXN, countYP, countYN := 0, 0, 0, 0, 0

				for j, r := range arr {
					for i, v := range r {
						if v == 0 && (i != x && j != y) { // we found an asteroid but not current one
							// get line func for current point and this one
							// y = kx + b
							k, b := lineFunc(float64(x), float64(y), float64(i), float64(j))

							fmt.Printf("Parent: (%d, %d) Child: (%d, %d) -- (%f, %f) \n", x, y, i, j, k, b)

							if _, ok := hmap[k]; ok {
								// if such key already exists than we append a new b
								hmap[k] = append(hmap[k], b)
							} else {
								hmap[k] = []float64{b}
							}
							count += len(hmap[k])
							//fmt.Printf("Map: %v \n", hmap)
						} else if v == 0 && ((i == x && j != y) || (i != x && j == y)) {
							if i == x && j > y { // above Y
								countYP = 1
							} else if i == x && j < y { // below Y
								countYN = 1
							} else if i > x && j == y { // right from X
								countXP = 1
							} else if i < x && j == y { // left from X
								countXN = 1
							}

							// get line func for current point and this one
							// y = kx + b
							k, b := lineFunc(float64(x), float64(y), float64(i), float64(j))
							fmt.Printf("Parent: (%d, %d) Child: (%d, %d) -- (%f, %f) \n", x, y, i, j, k, b)

							if _, ok := hmap[k]; ok {
								// if such key already exists than we append a new b
								hmap[k] = append(hmap[k], b)
							} else {
								hmap[k] = []float64{b}
							}
						}
					}
				}
				fmt.Printf("Map: %v \n", hmap)
				fmt.Printf("Count: %v \n", count)
				fmt.Printf("CountXP: %d CountXN: %d CountYP: %d CountYN: %d \n", countXP, countXN, countYP, countYN)
			}
		}
	}
}

func initialize(arr [][]int) [][]int {
	// Open file for reading
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	// X and Y axis lengths
	var x, y int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// read first row
		str := scanner.Text()
		// set X axis length
		x = len(str)
		// slice for X axis values
		var row []int

		// go through first row
		for _, char := range str {
			if char == '.' {
				row = append(row, -1)
			} else {
				row = append(row, 0)
			}
		}
		arr = append(arr, row)
	}
	// set Y axis length
	y = len(arr)

	fmt.Printf("X: %d  Y: %d \n", x, y)

	return arr
}

// Find line func between 2 points
// Returns y = (-A/b)x - C/b
func lineFunc(x1, y1, x2, y2 float64) (float64, float64) {
	a1 := y1 - y2
	b1 := x2 - x1
	c1 := x1*y2 - x2*y1

	//fmt.Printf("A1: %f B1: %f C1: %f \n", a1, b1, c1)

	var k float64
	if a1 == 0 || b1 == 0 {
		k = 0
	} else {
		k = -(a1 / b1)
	}

	var b float64
	if c1 == 0 || b1 == 0 {
		b = 0
	} else {
		b = -(c1 / b1)
	}

	//fmt.Printf("K: %f B: %f \n", k, b)

	return math.Round(k*100) / 100, math.Round(b*100) / 100
}
