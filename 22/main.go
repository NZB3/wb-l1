package main

import "fmt"

func main() {
	a := float64(1 << 65)
	b := float64(1 << 65)
	fmt.Println(a)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a + b)
	fmt.Println(a - b)
}
