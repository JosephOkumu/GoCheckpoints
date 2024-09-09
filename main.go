package main

import (
	"fmt"
	"os"
	"strconv"
)

func isPrime(nb int) bool {
	if nb == 1 {
		return false
	}

	for i := 2; i < nb; i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	if len(os.Args) != 2 {
		// In Go we do not use return in main. That's why we use os.Exit(1)
		os.Exit(1)
	}
	args := os.Args[1]
	num, err := strconv.Atoi(args)
	if err != nil {
		fmt.Printf("Error converting %s to type int\n", args)
	}
	fmt.Println(findPrevPrime(num))
	
}

func findPrevPrime(nb int) int {
	var temp []int
	for i := 1; i <= nb; i++ {
		if isPrime(i) {
			temp = append(temp, i)
		}		
	}
	
	length := len(temp)
	if length > 0 {
		return temp[length-1]
	}
	return 0
}
