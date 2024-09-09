package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	args := os.Args[1]

	num := 0
	sign := 1

	for i, v := range args {
		if i == 0 && v == '-' {
			sign = -1
		} else if i == 0 && v == '+' {
			sign = 1
		}
		if v >= 48 && v <= 57 {
			num = num*10 + int(v-48)
		}
	}
	newNum := sign * num
	tabMult(newNum)
}

func tabMult(newNum int) {
	for i := 1; i <= 9; i++ {
		result := newNum * i
		fmt.Printf("%d x %d = %d\n", i, newNum, result)
	}
}
