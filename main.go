package main

import "fmt"

func foldInt(f func(int, int) int, a []int, acc int) {
	result := acc

	for _, v := range a {
		result = f(result, v)
	}
	fmt.Println(result)
}

func Add(a, b int) int {
	return a + b
}

func Mul(a, b int) int {
	return a * b
}

func Sub(a, b int) int {
	return a - b
}

func main() {
	table := []int{1, 2, 3}
	ac := 93
	foldInt(Add, table, ac)
	foldInt(Mul, table, ac)
	foldInt(Sub, table, ac)
	fmt.Println()

	table = []int{0}
	foldInt(Add, table, ac)
	foldInt(Mul, table, ac)
	foldInt(Sub, table, ac)
}
