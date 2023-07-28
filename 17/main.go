package main

import (
	"fmt"
	"sort"

	"github.com/nzb3/randgen"
)

func main() {
	arr := randgen.GenArrayInt(100, 100)
	sort.Ints(arr)

	fmt.Println(arr)
	fmt.Println(BinSearch(arr, 10))
}

func BinSearch(arr []int, target int) int {
	left := 0
	right := len(arr) - 1

	for left <= right {
		mid := (left + right) / 2
		switch {
		case arr[mid] == target:
			return mid
		case arr[mid] < target:
			left = mid + 1
		case arr[mid] > target:
			right = mid - 1
		}
	}
	return -1
}
