package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"math/rand"
)

func main() {
	max := 20
	randomNum := rand.Intn(max) 

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		input := scanner.Text()
		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Not a number")
			return
		}
		if num == randomNum {
			fmt.Println("You are correct!")
			os.Exit(0)
		} else if num < randomNum {
			fmt.Println("Too low, try again")
		} else if num > randomNum {
			fmt.Println("Too high, try again")
		}
	}
}
