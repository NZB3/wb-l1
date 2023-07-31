package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	count int
}

func main() {
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}

	c := Counter{}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			mu.Lock()
			c.count++
			mu.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(c.count)
}
