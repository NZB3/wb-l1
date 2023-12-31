package main

import (
	"fmt"
)

func main() {
	set1 := []string{"cat", "cat", "dog", "cat", "tree"}

	m := make(map[string]struct{}, len(set1))
	for _, v := range set1 {
		m[v] = struct{}{}
	}

	var selfSet []string
	for k := range m {
		selfSet = append(selfSet, k)
	}
	fmt.Println(selfSet)
}
