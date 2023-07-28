package main

import (
	"fmt"
	"sync"
)

func sqNums(nums ...int) <-chan int {
	wg := sync.WaitGroup{}
	ch := make(chan int)
	go func() {

		for _, n := range nums {
			wg.Add(1)
			go func(n int) {
				ch <- n * n
				wg.Done()
			}(n)
		}
		wg.Wait()
		close(ch)
	}()
	return ch
}

func main() {
	sqNums := sqNums(2, 4, 6, 8, 10)
	sum := 0
	for n := range sqNums {
		sum += n
	}
	fmt.Println(sum)
}
