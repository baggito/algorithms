package main

import (
	"fmt"
)

// Function that reverse a given string
func reverse(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}

func main() {
	// given string
	var str = "lee)(t(c)o)((de)"
	// ( ( ( ) )( ) )

	// if we see "(" we do open++, if we see ")" we do open--
	open := 0
	leftToRight := ""
	rightToLeft := ""

	// Left to right traverse
	for _, c := range str {
		if string(c) == "(" {
			open++
		} else if string(c) == ")" {
			// In this step we have balanced string and current close ")" is redundant
			if open == 0 {
				continue // we ignore it
			}
			open--
		}
		// append parentheses (which doesn't break a balance) to leftToRight
		leftToRight += string(c)
	}

	fmt.Println(leftToRight)

	// Right to left traverse
	// In this step we have either balanced string or a string with redundant opens "("
	for i := len(leftToRight) - 1; i >= 0; i-- {
		if string(leftToRight[i]) == "(" {
			if open > 0 {
				open--
				continue
			}
		}

		rightToLeft += string(leftToRight[i])
	}

	fmt.Println(reverse(rightToLeft))
	fmt.Println("The end of algorithm")
}
