package main

import (
	"fmt"
)

func main() {
	fmt.Print("Enter string: ")
	var str string
	fmt.Scan(&str)
	fmt.Printf("Reversed string: %s\n", reverseString(str))
}

func reverseString(s string) string {
	runes := []rune(s)
	for i := 0; i < len(runes)/2; i++ {
		runes[i], runes[len(runes)-1] = runes[len(runes)-1], runes[i]
	}
	return string(runes)
}
