package main

import "fmt"

func swapBits(oct byte) byte {
	return (oct << 4) | (oct >> 4)
}

func main() {
	fmt.Println(swapBits(0x05)) // Output: 0x50
}
