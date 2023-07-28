package main

import (
	"fmt"
	"sync"
)

func main() {
	first()
	second()
	third()
}

// Mutex
func first() {
	m := make(map[int]int)
	var wg sync.WaitGroup
	var mu sync.Mutex

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			mu.Lock()
			m[i] = i
			mu.Unlock()
			wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Println(m)
}

// sync.Map
func second() {
	m := sync.Map{}
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			m.Store(i, i)
			wg.Done()
		}(i)
	}

	wg.Wait()
	m.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		return true
	})
}

// RWMutex
func third() {
	m := make(map[int]int)
	var wg sync.WaitGroup
	var mu sync.RWMutex

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			mu.Lock()
			m[i] = i
			mu.Unlock()
			wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Println(m)
}
