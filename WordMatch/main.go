package main

import (
    "fmt"
    "os"
)

// canRewrite checks if it is possible to rewrite the first string using characters from the second string while preserving the order.
func canRewrite(first, second string) bool {
    if len(first) == 0 || len(second) == 0 {
        return false
    }

    firstIndex, secondIndex := 0, 0
    for firstIndex < len(first) && secondIndex < len(second) {
        if first[firstIndex] == second[secondIndex] {
            firstIndex++
        }
        secondIndex++
    }
    // I am loving this code for wordmatch.
    return firstIndex == len(first)
}

func main() {
    // Check if the number of command-line arguments is 3 (including the program name)
    if len(os.Args) != 3 {
        return
    }

    first, second := os.Args[1], os.Args[2]
    // Check if the first string can be rewritten using characters from the second string
    if canRewrite(first, second) {
        fmt.Println(first)
    }
}