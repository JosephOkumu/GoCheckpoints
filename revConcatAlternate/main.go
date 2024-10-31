package main

import "fmt"

func RevConcatAlternate(slice1, slice2 []int) []int {
	result := []int{}
	var longer, shorter []int

	// Determine which slice is longer and which is shorter
	if len(slice1) > len(slice2) {
		longer, shorter = slice1, slice2
	} else {
		longer, shorter = slice2, slice1
	}

	// Add the remaining elements of the longer slice in reverse order
	for i := len(longer) - 1; i >= len(shorter); i-- {
		result = append(result, longer[i])
	}

	// Add elements from both slices in reverse order, alternating
	for i := len(shorter) - 1; i >= 0; i-- {
		result = append(result, slice1[i], slice2[i])
	}

	return result
}

func main() {
	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6}))        // Expected: [3 6 2 5 1 4]
	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6, 7, 8, 9})) // Expected: [9 8 7 3 6 2 5 1 4]
	fmt.Println(RevConcatAlternate([]int{1, 2, 3, 9, 8}, []int{4, 5}))      // Expected: [8 9 3 2 5 1 4]
	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{}))                // Expected: [3 2 1]
}
