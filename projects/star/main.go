package main

import (
	"fmt"
	"strings"
)

func addStars(s string) string {
	return strings.Join(strings.Split(s, ""), "*")
}

func main() {
	var input string
	fmt.Scan(&input)
	fmt.Println(addStars(input))
}
