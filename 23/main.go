package main

import (
	"fmt"
)

func main() {
	var idx int
	fmt.Scan(&idx)

	arr := []int{1, 2, 3, 4, 5}
	arr = remove(arr, idx)
	fmt.Println(arr)
}

func remove[T any](arr []T, idx int) []T {
	return append(arr[:idx], arr[idx+1:]...)
}
