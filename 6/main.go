package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	fmt.Println("Choose option to stop gorutine:")
	fmt.Println("1. Close done channel")
	fmt.Println("2. Call cancel function")
	fmt.Println("3. Stop main gorutine")
	fmt.Print("Your choice: ")
	var choice int
	fmt.Scan(&choice)

	ctx := context.Background()
	stop := make(chan bool)
	done := make(chan bool)

	ctx, cancel := context.WithCancel(ctx)
	go func() {
		fmt.Println("Gorutine is working 10 seconds...")
		time.Sleep(time.Second * 10)
		for {
			select {
			case <-ctx.Done():
				fmt.Println("context clanceled")
				done <- true
				return
			case <-stop:
				fmt.Println("stop channel closed")
				done <- true
				return
			}
		}
	}()

	switch choice {
	case 1:
		close(stop)
	case 2:
		cancel()
	case 3:
		fmt.Println("Waiting for 5 seconds...")
		time.Sleep(time.Second * 5)
		return
	}

	<-done
}
