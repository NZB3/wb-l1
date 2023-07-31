package main

import (
	"fmt"
	"sort"

	"github.com/nzb3/measure"
	"github.com/nzb3/randgen"
)

func main() {
	var size int
	fmt.Scan(&size)

	arr := randgen.GenArrayInt(size, size)
	sort.Ints(arr)

	fmt.Println(arr)

	var target int
	fmt.Scan(&target)

	measure.Measure(func() {
		fmt.Println(BinSearch(arr, target))
	})
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
