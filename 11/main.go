package main

import (
	"fmt"
	"sync"

	"github.com/nzb3/randgen"
)

func main() {
	mu := sync.RWMutex{}
	wg := sync.WaitGroup{}

	set1 := randgen.GenArrayInt(10, 10)
	set2 := randgen.GenArrayInt(10, 10)
	fmt.Println(set1)
	fmt.Println(set2)

	m := make(map[int][2]bool, len(set1))

	wg.Add(2)
	go func() {
		defer wg.Done()

		for _, n := range set1 {
			mu.RLock()
			if m[n] == [2]bool{true, true} {
				mu.RUnlock()
				continue
			}
			mu.RUnlock()

			mu.Lock()
			m[n] = [2]bool{true, m[n][1]}
			mu.Unlock()
		}

	}()
	go func() {
		defer wg.Done()

		for _, n := range set2 {
			mu.RLock()
			if m[n] == [2]bool{true, true} {
				mu.RUnlock()
				continue
			}
			mu.RUnlock()

			mu.Lock()
			m[n] = [2]bool{m[n][0], true}
			mu.Unlock()
		}

	}()

	wg.Wait()
	for k, v := range m {
		if v == [2]bool{true, true} {
			fmt.Println(k)
		}
	}
}
