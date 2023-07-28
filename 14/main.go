package main

import (
	"fmt"
	"reflect"
)

func main() {
	var data interface{}

	data = 1
	fmt.Println(reflect.TypeOf(data))

	data = "string"
	fmt.Println(reflect.TypeOf(data))

	data = make(chan int)
	fmt.Println(reflect.TypeOf(data))

	data = true
	fmt.Println(reflect.TypeOf(data))
}
