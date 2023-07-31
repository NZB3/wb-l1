package main

import "strings"

func reverseWords(s string) string {
	words := strings.Split(s, " ")
	for i := 0; i < len(words)/2; i++ {
		words[i], words[len(words)-1-i] = words[len(words)-1-i], words[i]
	}
	return strings.Join(words, " ")
}

func main() {
	str := "snow dog sun"
	str = reverseWords(str)
	println(str)
}
