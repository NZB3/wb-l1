package main

type Human struct {
	name string
}

type Action struct {
	Human
}

func (h *Human) SayName() {
	println(h.name)
}

func main() {
	a := new(Action)
	a.name = "nikolay"
	a.SayName()
}
