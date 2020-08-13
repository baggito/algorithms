package main

import (
	"fmt"
)

// N x N dashboard
const N int = 5

// Below arrays details all 8 possible movements for a knight.
// Don't change the sequence of below arrays
var row = []int{2, 1, -1, -2, -2, -1, 1, 2, 2}
var col = []int{1, 2, 2, 1, -1, -2, -2, -1, 1}

// Check if (x, y) is valid chess board coordinates
// Note that a knight cannot go out of the chessboard
func isValid(x, y int) bool {
	if x < 0 || y < 0 || x >= N || y >= N {
		return false
	}
	return true
}

func printArr(visited [N][N]int) {
	for _, v := range visited {
		fmt.Println(v)
	}
	fmt.Println("-------")
}

// Recursive function to perform the Knight's tour using backtracking
func knightTour(visited [N][N]int, x, y, pos int) {
	// mark current square as visited
	visited[x][y] = pos

	// if all squares are visited, print the solution
	if pos >= N*N {
		printArr(visited)
		// backtrack before returning
		visited[x][y] = 0
		return
	}

	// check for all 8 possible movements for a knight
	// and recur for each valid movement
	for k := 0; k < 8; k++ {
		// Get the new position of Knight from current
		// position on chessboard
		newX := x + row[k]
		newY := y + col[k]

		// if new position is a valid and not visited yet
		if isValid(newX, newY) && visited[newX][newY] == 0 {
			knightTour(visited, newX, newY, pos+1)
		}
	}

	// backtrack from current square and remove it from current path
	visited[x][y] = 0
}

func main() {

	// visited[][] serves two purpose -
	// 1. It keep track of squares involved the Knight's tour.
	// 2. It stores the order in which the squares are visited
	var visited = [N][N]int{}
	var pos = 1

	// start knight tour from corner square (0, 0)
	knightTour(visited, 0, 0, pos)

	fmt.Println("The end of algorithm")
}
