package main

import (
	// "github.com/01-edu/z01"
	"fmt"
)

func reduceInt(a []int, f func(int, int)int) {
	result := a[0]

	for _, v := range a[1:]{
		result = f(result, v)
	}
	fmt.Println(result)
}

func main(){
	mul := func(acc int, cur int) int{
		return acc * cur
	}
	sum := func(acc int, cur int) int {
		return acc + cur
	}
	div := func(acc int, cur int) int {
		return acc/ cur
	}
	as := []int{500, 2}

	reduceInt(as, mul)
	reduceInt(as, sum)
	reduceInt(as, div)

}