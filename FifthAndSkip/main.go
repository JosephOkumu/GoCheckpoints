package main

import (
    "fmt"
)

func FifthAndSkip(str string) string {
    if len(str) == 0 {
        return "\n"
    }    
    if len(str) < 5 {
        return "Invalid Input\n"
    }

    result := ""
    count := 0
    validCount := 0

    for i := 0; i < len(str); i++ {
        if str[i] == ' ' {
            continue
        }
        result += string(str[i])
        validCount++

        if validCount == 5 {
            // count how many groups of fives have been processed
            count++ 
            validCount = 0

            // Remove the next character (sixth)
            if i+1 < len(str) {
                i++ // skip the sixth character.
            }
            result += " "
        }
    }
    // Remove the last space if it exists
    if len(result) > 0 && result[len(result)-1] == ' ' {
        result = result[:len(result)-1]
    }
    return result + "\n"
}

func main() {
	fmt.Print(FifthAndSkip("abcdefghijklmnopqrstuwxyz"))
	fmt.Print(FifthAndSkip("This is a short sentence"))
	fmt.Print(FifthAndSkip("1234"))
}
