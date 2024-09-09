package main

import (
	"fmt"
	"os"
)

func piglatin(str string) string {
	vowels := "aeiouAEIOU"
	
	result := ""
	for i, v := range str {
		for _, ch := range vowels {
			if v == ch {
				result += str[i:] + str[:i] + "ay"
			}
			
		}
	}
	
	if len(result) != 0 {
		return result
	}
	return "No vowels."
	
}

func main() {
	if len(os.Args) != 2 {
		return
	}
	args1 := os.Args[1]

	fmt.Println(piglatin(args1))
}
