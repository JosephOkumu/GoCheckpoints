package main

import "fmt"

func atoiBase(s string, base string) int {
	// Check base validity
	if len(base) < 2 {
		return 0 // Base must have at least 2 characters
	}

	// Create a map to store unique characters in the base
	charSet := make(map[rune]bool)
	for _, char := range base {
		// If a character is +, -, or a duplicate
		if char == '+' || char == '-' || charSet[char] {
			// Base cannot contain '+', '-', or duplicate characters
			return 0 
		}
		charSet[char] = true
	}

	// Create a map for quick digit lookup
	digitValue := make(map[rune]int)
	for i, digit := range base {
		digitValue[digit] = i
	}

	result := 0
	baseLen := len(base)

	// Iterate over each character in the input string
	for _, char := range s {
		value, exists := digitValue[char]
		if !exists {
			return 0 // Invalid character in the string
		}
		result = result*baseLen + value
	}

	return result
}

func main() {
	fmt.Println(atoiBase("125", "0123456789"))
	fmt.Println(atoiBase("1111101", "01"))
	fmt.Println(atoiBase("7D", "0123456789ABCDEF"))
	fmt.Println(atoiBase("uoi", "choumi"))
	fmt.Println(atoiBase("bbbbbab", "-ab"))
}
