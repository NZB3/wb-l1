package main

import "fmt"

func makeChan(nums ...int) <-chan int {
	ch := make(chan int)

	go func() {
		for _, n := range nums {
			ch <- n
		}
		close(ch)
	}()

	return ch
}

func square(n int) <-chan int {
	ch := make(chan int)
	go func() {
		ch <- n * n
	}()
	return ch
}

func sqNums(nums <-chan int) <-chan int {
	ch := make(chan int)
	go func() {
		for n := range nums {
			ch <- <-square(n)
		}
		close(ch)
	}()
	return ch
}

func main() {
	nums := makeChan(2, 4, 6, 8, 10)
	sqNums := sqNums(nums)
	sum := 0
	for n := range sqNums {
		sum += n
	}
	fmt.Println(sum)
}
