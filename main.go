package main

import "fmt"

func repeatAlpha(str string) string {
	result := ""
	for _, v := range str {
		if (v >= 'a' && v <= 'z') {
			index := int(v) - int('a') + 1
			for i := 0; i < index; i++ {
				result += string(v)
			}
		} else if (v >= 'A' && v <= 'Z') {
			index := int(v) - int('A') + 1
			for i := 0; i < index; i++ {
				result += string(v)
			}
		} else {
			result += string(v)
		}
	}
	return result
}

func main() {
	fmt.Println(repeatAlpha("abc"))         // Expected output: abbccc
	fmt.Println(repeatAlpha("Choumi."))     // Expected output: CChhhhhhoooooooooouuuuuuuuummmmmmmmmmiiiiiiii.
	fmt.Println(repeatAlpha(""))            // Expected output: (empty string)
	fmt.Println(repeatAlpha("abacadaba 01!")) // Expected output: abbccacaddabbbb 01!
}
