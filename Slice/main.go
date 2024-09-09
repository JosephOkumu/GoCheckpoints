package main

import "fmt"

func slice(a []string, nbrs... int) []string {
	if len(nbrs) == 1 {
		if nbrs[0] < 0 {
			nbrs[0] = -nbrs[0]
			return a[nbrs[0]-1:]
		}
		return a[nbrs[0]:]
	}
	if len(nbrs) == 2 {
		if (nbrs[0] > 0 && nbrs[1] > 0) && (nbrs[1] > nbrs[0]) {
			return a[nbrs[0]:nbrs[1]]
		} else if (nbrs[0] < 0 && nbrs[1] < 0) && (nbrs[0] < nbrs[1]) {
			result := -(nbrs[0] + nbrs[1])
			return a[result : result+1]
		}
	}
	return nil
}

func main() {
	a := []string{"coding", "algorithm", "ascii", "package", "golang"}
	fmt.Printf("%#v\n", slice(a, 1))
	fmt.Printf("%#v\n", slice(a, 2, 4))
	fmt.Printf("%#v\n", slice(a, -3))
	fmt.Printf("%#v\n", slice(a, -2, -1))
	fmt.Printf("%#v\n", slice(a, 2, 0))
}
