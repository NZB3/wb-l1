package main

import "fmt"

func main() {
	fmt.Print("Enter string: ")
	var str string
	fmt.Scan(&str)

	uniq := make(map[rune]bool, len(str))

	for _, r := range str {
		uniq[r] = true
	}

	if len(uniq) == len(str) {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}
}
