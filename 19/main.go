package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Print("Enter string: ")
	var str string
	fmt.Scan(&str)

	wg := sync.WaitGroup{}

	wg.Add(2)
	go func() {
		fmt.Println("Reversed string 1:", reverse1(str))
		wg.Done()
	}()
	go func() {
		fmt.Println("Reversed string 2:", reverse2(str))
		wg.Done()
	}()

	wg.Wait()

}

func reverse2(str string) string {
	if len(str) == 0 {
		return ""
	}
	return reverse2(str[1:]) + string(str[0])
}

func reverse1(str string) string {
	sp := 0
	ep := len(str) - 1

	for sp < ep {
		str = swap(str, sp, ep)
		sp++
		ep--
	}

	return str
}

func swap(str string, sPointer, ePointer int) string {
	strBytes := []byte(str)
	strBytes[sPointer], strBytes[ePointer] = strBytes[ePointer], strBytes[sPointer]
	return string(strBytes)
}
