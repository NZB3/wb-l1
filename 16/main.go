package main

import (
	"fmt"

	"github.com/nzb3/randgen"
)

func main() {
	arr := randgen.GenArrayInt(100000, 100000)
	arr = QuickSort(arr)

	fmt.Println(arr)
}

func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivot := arr[int(len(arr)/2)]
	left := make([]int, 0)
	right := make([]int, 0)
	equal := make([]int, 0)
	for _, num := range arr {
		if num < pivot {
			left = append(left, num)
		} else if num > pivot {
			right = append(right, num)
		} else {
			equal = append(equal, num)
		}
	}

	left = QuickSort(left)
	right = QuickSort(right)

	return append(append(left, equal...), right...)
}
