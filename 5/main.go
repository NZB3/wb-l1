package main

import (
	"fmt"
	"time"
)

func startWriter(ch chan<- int, timer *time.Timer) {
	for i := 0; ; i++ {
		select {
		case <-timer.C:
			close(ch)
			fmt.Println("Sending data is stoped")
			return
		default:
			ch <- i
		}
	}
}

func main() {
	var secCount int
	fmt.Scan(&secCount)

	start := time.Now()

	ch := make(chan int)
	timer := time.NewTimer(time.Second * time.Duration(secCount))
	go startWriter(ch, timer)

	for data := range ch {
		fmt.Println(data)
	}

	fmt.Printf("Time: %s's\n", time.Since(start))
}
