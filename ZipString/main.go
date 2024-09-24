package main

import (
	"fmt"
)

func zipString(s string) string {
	var result string
	var prevChar rune
	var count int

	for _, char := range s {
		if prevChar == 0 || char != prevChar {
			if prevChar != 0 {
				result += fmt.Sprintf("%d%c", count, prevChar)
			}
			prevChar = char
			count = 1
		} else {
			count++
		}
	}

	if prevChar != 0 {
		result += fmt.Sprintf("%d%c", count, prevChar)
	}

	return result

}

func main() {
	fmt.Println(zipString("YouuungFellllas"))                                   // Output: 1Y1o3u1n1g1F1e4l1a1s
	fmt.Println(zipString("Thee quuick browwn fox juumps over the laaazy dog")) // Output: 1T1h2e1 1q2u1i1c1k1 1b1r1o2w1n1 1f1o1x1 1j2u1m1p1s1 1o1v1e1r1 1t1h1e1 1l3a1z1y1 1d1o1g
	fmt.Println(zipString("Helloo Therre!"))                                    // Output: 1H1e2l2o1 1T1h1e2r1e1!
}
