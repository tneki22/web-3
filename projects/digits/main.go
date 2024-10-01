package main

import (
	"fmt"
)

func findMaxDigit(s string) rune {
	maxDigit := '0'

	for _, digit := range s {
		if digit > maxDigit {
			maxDigit = digit
		}
	}

	return maxDigit
}

func main() {
	var input string
	fmt.Scan(&input)
	fmt.Printf("%c\n", findMaxDigit(input))
}
