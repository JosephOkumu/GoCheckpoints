package main

import "fmt"

func reverseBits(oct byte) byte {
	var result byte
	for i := 0; i < 8; i++ {
		result = (result << 1) | (oct & 1)
		oct >>= 1
	}
	return result
}

func main() {
	example := byte(0x26) // 0010 0110
	reversed := reverseBits(example)
	fmt.Printf("%08b || / %08b\n", example, reversed) // Output: 00100110 || / 01100100
}
