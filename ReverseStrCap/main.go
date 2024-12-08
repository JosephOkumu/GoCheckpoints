package main

import (
	"fmt"
	"os"
)

func reverseStrCap(str string) string {
	currentWord := ""
	result := ""

	for i, char := range str {
		if char >= 65 && char <= 90 {
			char += 32
		}
		if char != ' ' {
			currentWord += string(char)
		}
		if char == ' ' || i == len(str)-1 {
			if len(currentWord) > 0 {
				lastChar := currentWord[len(currentWord)-1]
				if lastChar >= 97 && lastChar <= 122{
				lastChar -= 32
				}
				currentWord = currentWord[:len(currentWord)-1] + string(lastChar)
				result += currentWord

				if char == ' ' {
					result += " "
					currentWord = ""
				}
			}
			
		}
	}
	return result
}

func main() {
	if len(os.Args) == 1 {
		return
	}
	args := os.Args[1:]

	for _, arg := range args {
		fmt.Println(reverseStrCap(arg))
	}
}
