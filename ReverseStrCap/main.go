package main

import "os"

// revstrcap takes a string and reverses the capitalization of the last letter of each word
// while converting all other letters to lowercase.
func reverseStrCap(str string) string {
	var result string
	var currentWord string // Variable to hold the current word being processed

	for i, v := range str {
		// Convert uppercase letters to lowercase
		if v >= 'A' && v <= 'Z' {
			v += 32 
		}
		if v != ' ' {
			currentWord += string(v)

		}

		// Check if the character is a space or if it's the last character in the string
		if v == ' ' || i == len(str)-1 {
			// Process the current word if it's not empty
			

			lastChar := currentWord[len(currentWord)-1]
			// Convert the last character to uppercase
			lastChar -= 32
			// Append the current word (except the last character) and the uppercase last character to the result
			result += currentWord[:len(currentWord)-1] + string(lastChar)
			

			// If the character is a space, add a space to the result and reset currentWord
			if v == ' ' {
				result += " "
				currentWord = "" 
			}
		}
	}

	return result 
}

func main() {
	if len(os.Args) < 2 {
		return
	}
	args := os.Args[1:]

	for _, arg := range args {
		os.Stdout.WriteString(reverseStrCap(arg) + "\n")
	}
}
