package main

import (
	"fmt"
	"time"

	"github.com/nzb3/randgen"
)

func main() {
	var size int
	fmt.Scan(&size)
	arr := randgen.GenArrayInt(size, size)

	start := time.Now()
	arr = QuickSort(arr)
	time := time.Since(start)

	fmt.Println(arr)
	fmt.Println(time)

}

func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivot := arr[int(len(arr)/2)]
	left := make([]int, 0, cap(arr)/2)
	right := make([]int, 0, cap(arr)/2)
	equal := make([]int, 0, 1)

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
