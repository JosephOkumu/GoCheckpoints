package main

import (
	"fmt"
	"os"
	"strconv"
)

func isPowerOf2(num int) bool {
	if num == 1 {
		return true
	} else if num < 1 {
		return false
	}
	for num > 1 {		
		if num%2 != 0 {
			return false			
		}
		num /= 2
	}
	return true
}

func main() {
	if len(os.Args) != 2 {
		return
	}
	str := os.Args[1]
	num, _ := strconv.Atoi(str)
	fmt.Println(isPowerOf2(num))
}
