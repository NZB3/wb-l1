package main

import (
	"fmt"

	"github.com/nzb3/randgen"
)

func makeChan(nums []int) <-chan int {
	ch := make(chan int)

	go func() {
		for _, n := range nums {
			ch <- n
		}
		close(ch)
	}()

	return ch
}

func main() {
	arr := randgen.GenArrayInt(10, 10)
	fmt.Println(arr)
	ch := makeChan(arr)
	sqCh := make(chan int)

	go func() {
		for n := range ch {
			sqCh <- n * n
		}
		close(sqCh)
	}()

	for n := range sqCh {
		fmt.Println(n)
	}
}
