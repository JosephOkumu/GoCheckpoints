package main

import (
	"os"
	"fmt"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("options: abcdefghijklmnopqrstuvwxyz")
		return
	}
	args := os.Args[1:]
	
	options := 0
	for _, arg := range args {
		if arg[0] != '-' || len(arg) == 1 {
			fmt.Println("Invalid Option")
			return
		}
		if arg[1] == 'h' {
			fmt.Println("options: abcdefghijklmnopqrstuvwxyz")
			return
		}
		for _, char := range arg[1:] {
			if char <  97 || char > 122 {
				fmt.Println("Invalid Option")
				return
			}
			options |= 1 << (char - 'a')
		}
	}
	result := fmt.Sprintf("%b", options)
	padding := 32 - len(result)

	for i := 0; i < padding; i++ {
		result = "0" + result		
	}
	
	result = result[:8] + " " + result[8:16] + " " + result[16:24] + " " + result[24:] 
	fmt.Println(result)
}
