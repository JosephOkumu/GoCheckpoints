package main

import (
	"fmt"
	"os"
)

// PrintMemory takes an array of 10 bytes and prints the memory representation
func PrintMemory(arr [10]byte) {
	// Print the hexadecimal values of the bytes
	for i := 0; i < 10; i++ {
		// Add a newline after every 4 bytes for readability
		if i > 0 && i%4 == 0 {
			os.Stdout.WriteString("\n")
		}
		// Format the byte as a hexadecimal string
		if i < 9 {
			os.Stdout.WriteString(fmt.Sprintf("%02x ", arr[i]))
		} else {
			os.Stdout.WriteString(fmt.Sprintf("%02x", arr[i]))
		}
	}
	os.Stdout.WriteString("\n")

	// Print the printable characters represented by the bytes
	for _, b := range arr {
		// Check if the byte represents a printable character (ASCII code between 32 and 126)
		if b >= 32 && b <= 126 {
			os.Stdout.WriteString(string(b))
		} else {
			// If the byte is not printable, print a dot
			os.Stdout.WriteString(".")
		}
	}
	os.Stdout.WriteString("\n")
}

func main() {
	// Create an array of 10 bytes and pass it to the PrintMemory function
	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
}
