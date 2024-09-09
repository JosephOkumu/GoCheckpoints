package main

import (
	"fmt"
	"strconv"
)

func NotDecimal(dec string) string {
	// Return a newline if the input is empty
	if dec == "" {
		return "\n"
	}

	// Variables to store the integer part and fractional part
	var integerPart, fractionalPart string
	dotFound := false

	// Iterate over each character in the string
	for _, char := range dec{
		if char == '.' {
			dotFound = true
			continue
		}

		if dotFound {
			fractionalPart += string(char)
		} else {
			integerPart += string(char)
		}
	}

	// If no decimal point was found, return the original string followed by a newline
	if !dotFound {
		return dec + "\n"
	}

	// If the fractional part is "0", return the integer part followed by a newline
	if fractionalPart == "0" {
		return integerPart + "\n"
	}

	// Convert the integer part and fractional part to integers
	_, err1 := strconv.Atoi(integerPart)
	_, err2 := strconv.Atoi(fractionalPart)

	// If either conversion fails, return the original string followed by a newline
	if err1 != nil || err2 != nil {
		return dec + "\n"
	}

	// Return the concatenated integer and fractional parts without the decimal point
	return integerPart + fractionalPart + "\n"
}

func main() {
	fmt.Print(NotDecimal("0.1"))        // Expected: "01\n"
	fmt.Print(NotDecimal("174.2"))      // Expected: "1742\n"
	fmt.Print(NotDecimal("0.1255"))     // Expected: "01255\n"
	fmt.Print(NotDecimal("1.20525856")) // Expected: "120525856\n"
	fmt.Print(NotDecimal("-0.0f00d00")) // Expected: "-0.0f00d00\n"
	fmt.Print(NotDecimal(""))           // Expected: "\n"
	fmt.Print(NotDecimal("-19.525856")) // Expected: "-19525856\n"
	fmt.Print(NotDecimal("1952"))       // Expected: "1952\n"
}
