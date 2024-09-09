package main

import (
	"fmt"
	"os"
)

func mirrorAlpha(str string) string {
	result := ""
	for _, v := range str {
		if v >= 'A' && v <= 'Z' {
			distance := v - 'A'
			v = 'Z' - distance
		} else if v >= 'a' && v <= 'z' {
			distance := v - 'a'
			v = 'z' - distance
		}
		result += string(v)
	}
	return result
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println()
		return
	}
	inputStr := os.Args[1]
	fmt.Println(mirrorAlpha(inputStr))
}
