package main

import (
	"fmt"
)

func Chunk(slice []int, size int) {
	// Check if the size is 0, if so, print a newline and return
	if size == 0 {
		fmt.Println()
		return
	}

	// Initialize a 2D slice to store the chunked sub-slices
	var result [][]int

	// Iterate over the input slice in chunks of the specified size
	for i := 0; i < len(slice); i += size {
		// Determine the end index of the current chunk
		end := i + size
		if end > len(slice) {
			end = len(slice)
		}

		// Create a temporary slice to hold the current chunk
		temp := []int{}
		// Append the elements from the original slice to the temporary slice
		temp = append(temp, slice[i:end]...)
		// Append the temporary slice to the result slice
		result = append(result, temp)
	}

	// Print the resulting chunked sub-slices
	fmt.Println(result)
}

func main() {
	// Test cases with different inputs
	Chunk([]int{}, 10)                          // Output: []
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 0)     // Output: 
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 3)     // Output: [[0 1 2] [3 4 5] [6 7]]
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 5)     // Output: [[0 1 2 3 4] [5 6 7]]
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 4)     // Output: [[0 1 2 3] [4 5 6 7]]
}
