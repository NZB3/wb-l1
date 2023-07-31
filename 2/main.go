package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := []int{2, 4, 6, 8, 10}
	first(arr)
	second(arr)
}

func first(arr []int) {
	out := make(chan int)
	go func() {
		for _, n := range arr {
			out <- n
		}
		close(out)
	}()

	for n := range out {
		fmt.Println(n * n)
	}
}

func second(arr []int) {
	var wg sync.WaitGroup

	wg.Add(len(arr))
	for i := range arr {
		go func(i int) {
			arr[i] = arr[i] * arr[i]
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println(arr)
}