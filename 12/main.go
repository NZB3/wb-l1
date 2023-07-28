package main

import (
	"fmt"
)

func main() {
	set1 := []string{"cat", "cat", "dog", "cat", "tree"}

	m := make(map[string]int, len(set1))
	for _, v := range set1 {
		m[v]++
	}

	var selfSet []string
	for k := range m {
		selfSet = append(selfSet, k)
	}
	fmt.Println(selfSet)
}
