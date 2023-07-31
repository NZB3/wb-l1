package main

import (
	"fmt"
)

func toggleBit(n int64, pos uint) int64 {
	// сдвигаем 1 на pos позиций влево
	mask := int64(1 << (pos - 1))
	// xor инвертирует бит по маске
	return mask ^ n
}

func main() {
	var num int64
	fmt.Print("Enter number: ")
	fmt.Scan(&num)

	var idx uint
	fmt.Print("Enter bit index: ")
	fmt.Scan(&idx)
	fmt.Println()

	fmt.Printf("Before:\t%b\n", num)
	num = toggleBit(num, idx)
	fmt.Printf("After:\t%b\n", num)
}
