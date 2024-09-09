package main

import (
	"os"

	"github.com/01-edu/z01"
)

func searchReplace(str, ltr1, ltr2 string) string {
	newStr := ""
	for _, v := range str {
		for _, k := range ltr1 {
			for _, l := range ltr2 {
				if v == k {
					v = l
				}
			}
		}
		newStr += string(v)
	}
	return newStr
}

func main() {
	if len(os.Args) != 4 {
		return
	}

	str := os.Args[1]
	ltr1 := os.Args[2]
	ltr2 := os.Args[3]

	word := searchReplace(str, ltr1, ltr2)

	for _, v := range word {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}
