package main

import "fmt"


func ConcatAlternate(slice1, slice2 []int) []int {
    result := []int{}
    smaller := []int{}
    bigger := []int{}

    if len(slice1) > len(slice2) {
        bigger = slice1
        smaller = slice2
    }else if len(slice2) > len(slice1) {
        bigger = slice2
        smaller = slice1
    }else {
        bigger = slice1
        smaller = slice2
    }

    for i := 0; i < len(smaller); i++ {
        result = append(result, bigger[i], smaller[i])
    }

    for i := len(smaller); i < len(bigger); i++ {
        result = append(result, bigger[i])
    }
    return result
}

func main() {
    fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6}))
    fmt.Println(ConcatAlternate([]int{2, 4, 6, 8, 10}, []int{1, 3, 5, 7, 9, 11}))
    fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6, 7, 8, 9}))
    fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{}))
}