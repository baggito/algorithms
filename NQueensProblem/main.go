package main

import (
	"fmt"
)

const N = 8

func isSafe(mat [N][N]string, r, c int) bool {
	// return false if two queens share the same column
	for i := 0; i < r; i++ {
		if mat[i][c] == "Q" {
			return false
		}
	}

	// return false if two queens share the same \ diagonal
	for i, j := r, c; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if mat[i][j] == "Q" {
			return false
		}
	}

	// return false if two queens share the same / diagonal
	for i, j := r, c; i >= 0 && j < N; i, j = i-1, j+1 {
		if mat[i][j] == "Q" {
			return false
		}
	}
	return true
}

func nQueen(mat [N][N]string, r int) {
	// if N queens are placed successfully, print the solution
	if r == N {
		for i := 0; i < N; i++ {
			for j := 0; j < N; j++ {
				fmt.Print(mat[i][j], " ")
			}
			fmt.Println()
		}
		fmt.Println()
		return
	}

	// place Queen at every square in current row r
	// and recur for each valid movement
	for i := 0; i < N; i++ {
		// if no two queens threaten each other
		if isSafe(mat, r, i) {
			// place queen on current square
			mat[r][i] = "Q"

			// recur for next row
			nQueen(mat, r+1)

			// backtrack and remove queen from current square
			mat[r][i] = "-"
		}
	}
}

func main() {

	// mat[][] keeps track of position of Queens in
	// the current configuration
	mat := [N][N]string{}

	// initialize mat[][] by '-'
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			mat[i][j] = "-"
		}
	}

	nQueen(mat, 0)

	fmt.Println("The end of algorithm")
}
