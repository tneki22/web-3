package main

import (
	"fmt"
	"strconv"
)

func squareDigits(num int) int {
	strNum := strconv.Itoa(num)
	result := ""

	for _, char := range strNum {
		digit, _ := strconv.Atoi(string(char))
		squared := digit * digit
		result += strconv.Itoa(squared)
	}

	finalResult, _ := strconv.Atoi(result)
	return finalResult
}

func main() {
	var input int
	fmt.Scan(&input)
	fmt.Println(squareDigits(input))
}
