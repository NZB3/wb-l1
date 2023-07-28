package main

import (
	"fmt"
	"sync"

	"github.com/nzb3/randgen"
)

func main() {
	mu := sync.Mutex{}
	wg := sync.WaitGroup{}

	set1 := randgen.GenArrayInt(10, 10)
	set2 := randgen.GenArrayInt(10, 10)
	fmt.Println(set1)
	fmt.Println(set2)

	m := make(map[int][2]bool, len(set1))
	wg.Add(2)
	go func() {
		for _, n := range set1 {
			if [2]bool{true, true} == m[n] {
				continue
			}
			mu.Lock()
			m[n] = [2]bool{true, m[n][1]}
			mu.Unlock()
		}
		wg.Done()
	}()

	go func() {
		for _, n := range set2 {
			if [2]bool{true, true} == m[n] {
				continue
			}
			mu.Lock()
			m[n] = [2]bool{m[n][0], true}
			mu.Unlock()
		}
		wg.Done()
	}()

	wg.Wait()
	for k, v := range m {
		if v == [2]bool{true, true} {
			fmt.Println(k)
		}
	}
}
