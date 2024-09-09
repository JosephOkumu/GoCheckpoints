package main

import (
	"fmt"
	"os"
)

// inter is a function that takes two strings, str1 and str2, and returns a new string
// that contains characters that appear in both string, in the order they appear in the first one without doulbes
func inter(str1, str2 string) string {
	// Initialize an empty string to store the result.
	var newStr string

	// Create a map to keep track of characters seen in str2.
	seen := map[rune]bool{}

	// Iterate over each character in str2.
	for _, char := range str2 {
		// Mark the character as seen.
		seen[char] = true
	}

	// Iterate over each character in str1.
	for _, v := range str1 {
		// If the character is seen in str2, add it to the result.
		if seen[v] {
			newStr += string(v)
			// Remove the character from the seen map.
			delete(seen, v)
		}
	}

	// Return the new string.
	return newStr
}

func main() {
	// Check if the correct number of command-line arguments are provided.
	if len(os.Args) != 3 {
		return
	}

	// Get the command-line arguments.
	str1 := os.Args[1]
	str2 := os.Args[2]

	// Call the inter function with the provided strings.
	fmt.Println(inter(str1, str2))
}
