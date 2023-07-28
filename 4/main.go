package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
)

type worker struct {
	id int
}

func (w *worker) work(dataChannel <-chan int, ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d done work\n", w.id)
			return
		default:
			for data := range dataChannel {
				fmt.Printf("Worker %d: %d\n", w.id, data)
			}
		}
	}

}

func spawnWorker(id int, dataChannel <-chan int, ctx context.Context, wg *sync.WaitGroup) {
	w := worker{
		id: id,
	}
	go w.work(dataChannel, ctx, wg)
}

func dataSender(ctx context.Context) <-chan int {
	ch := make(chan int)
	go func() {
		for i := 0; ; i++ {
			select {
			case <-ctx.Done():
				close(ch)
				fmt.Println("Sending data is stoped")
				return
			default:
				ch <- i
			}
		}

	}()
	return ch
}

func main() {
	ctx := context.Background()
	wg := &sync.WaitGroup{}

	var n int
	fmt.Print("Enter workers count: ")
	if _, err := fmt.Scan(&n); err != nil {
		fmt.Println(err)
	}

	ctx, cancel := context.WithCancel(ctx)
	dataChanel := dataSender(ctx)

	wg.Add(n)
	for i := 1; i <= n; i++ {
		spawnWorker(i, dataChanel, ctx, wg)
	}

	sig := make(chan os.Signal, 1)

	signal.Notify(sig, os.Interrupt)

	go func() {
		sig := <-sig
		fmt.Println()
		fmt.Println(sig)
		cancel()
	}()

	wg.Wait()
	fmt.Println("Program canceled")
}
