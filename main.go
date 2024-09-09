package main

import "fmt"

func compare(a, b string) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	} else {
		return 0
	}


	// This code below should also work the same way.
	// runesA := []rune(a)
	// runesB := []rune(b)
	// i := 0
	// j := 0

	// for i < len(runesA) && j < len(runesB) {
	// 	if runesA[i] < runesB[j] {
	// 		return -1
	// 	}else if runesA[i] > runesB[j] {
	// 		return 1
	// 	}
	// 	i++
	// 	j++
	// }	
	// if len(runesA) < len(runesB) {
	// 	return -1
	// }else if len(runesA) > len(runesB){
	// 	return 1
	// }
	// return 0

}

func main() {
	fmt.Println(compare("Hello!", "Hello!"))
	fmt.Println(compare("Salut!", "lut!"))
	fmt.Println(compare("Ola!", "Ol"))
}
