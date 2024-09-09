package main

import (
	"fmt"
	"os"
)

func union(str1, str2 string) string {
	newStr := ""
	full := str1 + str2
	seen := make(map[rune]bool)

	for _, char := range full {
		seen[char] = true
	}

	for _, char := range full {
		if seen[char] {
			newStr += string(char)
			delete(seen, char)
		}
	}
	return newStr
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println()
		os.Exit(1)
	}

	str1 := os.Args[1]
	str2 := os.Args[2]

	fmt.Println(union(str1, str2))
}
