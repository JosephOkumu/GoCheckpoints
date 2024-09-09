package main

import (
	"fmt"
	"os"
	"strconv"
)

func toHex(n int) string {
	if n == 0 {
		return "0"
	}
	hexChars := "0123456789abcdef"
	result := ""

	for n > 0 {
		remainder := n % 16
		result = string(hexChars[remainder]) + result
		n /= 16
	}
	return result
}

func main() {
	if len(os.Args) != 2 {
		return
	}
	input := os.Args[1]
	num, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Error")
		return
	}

	fmt.Println(toHex(num))
}
