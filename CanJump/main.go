package main

import "fmt"

func canJump(slice []uint) bool {
	if len(slice) == 0 {
		return false
	} else if len(slice) == 1 {
		return true
	}
	for i := 0; i < len(slice); i++ {
		result := uint(i) + slice[i]
		if result == uint(len(slice)-1) {
			return true
		}
	}
	return false
}

func main() {
	input1 := []uint{2, 3, 1, 1, 4}
	fmt.Println(canJump(input1))

	input2 := []uint{3, 2, 1, 0, 4}
	fmt.Println(canJump(input2))

	input3 := []uint{0}
	fmt.Println(canJump(input3))
}
