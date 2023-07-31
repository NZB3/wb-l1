package main

import "fmt"

type Type1 interface {
	Method1()
}

type Type2 struct {
}

func (t Type2) Method2() {
	fmt.Print("Method2")
}

type Adapter struct {
	Type2
}

func (a Adapter) Method1() {
	a.Method2()
}

func main() {
	t1 := Type1(Adapter{})
	t1.Method1()
}
