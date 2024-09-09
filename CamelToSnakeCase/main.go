package main

import (
	"fmt"
)

func CamelToSnakeCase(s string) string {
	// Check if the string is empty
	if s == "" {
		return ""
	}

	// Initialize the result string
	var result string

	// Previous character tracker
	var prevIsUpper bool

	// Iterate over the string
	for i := 0; i < len(s); i++ {
		c := s[i]

		// Check if the character is an uppercase letter
		if c >= 'A' && c <= 'Z' {
			// If it's not the first character, prepend an underscore
			if i > 0 {
				result += "_"
			}
			// Convert the uppercase letter to lowercase
			c += 'a' - 'A'

			// Check if the previous character was also uppercase
			if prevIsUpper {
				return s
			}
			prevIsUpper = true
		} else {
			// Reset the tracker if the character is not uppercase
			prevIsUpper = false
		}

		// Append the character to the result
		result += string(c)
	}

	// Check if the string ends with an uppercase letter
	if prevIsUpper {
		return s
	}

	return result
}

func main() {
	fmt.Println(CamelToSnakeCase("HelloWorld"))        // hello_world
	fmt.Println(CamelToSnakeCase("helloWorld"))        // hello_world
	fmt.Println(CamelToSnakeCase("camelCase"))         // camel_case
	fmt.Println(CamelToSnakeCase("CAMELtoSnackCASE"))  // CAMELtoSnackCASE
	fmt.Println(CamelToSnakeCase("camelToSnakeCase"))  // camel_to_snake_case
	fmt.Println(CamelToSnakeCase("hey2"))              // hey2
}
