package main

import (
	"fmt"
	"os"
)

func expandStr(str string) string {
	word := ""
	result := ""
	for i, char := range str {
		if char != ' ' {
			word += string(char)
		}
		if char == ' ' || i == len(str)-1{
			if len(word) > 0 {
				word += "   "
				result += word
				word = ""
			}
		}
	}

	for i := 0; i < len(result); i++ {
		if result[len(result)-1] == ' '{
			result = result[:len(result)-1]
		}
	}
	return result
}

func main() {
	if len(os.Args) != 2 {
		return
	}
	args := os.Args[1]
	fmt.Println(expandStr(args))
}
