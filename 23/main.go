package main

import (
	"fmt"
)

func main() {
	var idx int
	fmt.Scan(&idx)

	arr := []int{1, 2, 3, 4, 5}
	fmt.Printf("Before: %v\n", arr)
	arr = remove2(arr, idx)
	fmt.Printf("After deliting %d: %v", idx, arr)
	arr = remove1(arr, idx)

	fmt.Println(arr)
}

func remove1[T any](arr []T, idx int) []T {
	return append(arr[:idx], arr[idx+1:]...)
}

func remove2[T any](slice []T, idx int) []T {
	slice[idx] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}

// Можно написать
// hdr := (*reflect.SliceHeader)(unsafe.Pointer(&slice))
// data := *(*[5]int)(unsafe.Pointer(hdr.Data))
// fmt.Print(data)
// У вас получится [1 2 5 4 5]
