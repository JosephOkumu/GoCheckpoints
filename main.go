package main

import "fmt"

func sortwordArray(slice []string) []string {
	for i := 0; i < len(slice)-1; i++ {
		for j := i + 1; j < len(slice); j++ {
			if slice[j] < slice[i] {
				slice[i], slice[j] = slice[j], slice[i]
			}
		}
	}
	return slice
}

func main() {
	result := []string{"a", "A", "1", "b", "B", "2", "c", "C", "3"}
	sortwordArray(result)

	fmt.Println(result)
}
