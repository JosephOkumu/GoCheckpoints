package main

import (
	"fmt"
)

func isCapitalized(s string) bool {
	if s == "" {
		return false
	}
	words := []string{}
	result := ""

	for _, ch := range s {
		if ch != ' ' {
			result += string(ch)
		}else if ch == ' ' && result != "" {
			words = append(words, result)
			result = ""
		}
	}
	words = append(words, result)

	for _, word := range words {
		firstChar := word[0]

		if firstChar >= 'a' && firstChar <= 'z' {
			return false
		}
	}
	return true
}


func main() {
	fmt.Println(isCapitalized("Hello! How are you?"))
	fmt.Println(isCapitalized("Hello How Are You"))
	fmt.Println(isCapitalized("Whats 4this 100K?"))
	fmt.Println(isCapitalized("Whatsthis4"))
	fmt.Println(isCapitalized("!!!!Whatsthis4"))
	fmt.Println(isCapitalized(""))
}

