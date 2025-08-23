package main

import "fmt"

type Human struct {
	Name string
	Age  int
}

func (h *Human) Talk(s string) {
	fmt.Println(s)
}

func (h *Human) GrowUp() {
	h.Age++
}

func (h *Human) GetAge() int {
	return h.Age
}

type Action struct {
	Human
}

func (a *Action) Jump() {
	fmt.Println("Jump")
}

func main() {
	Danila := Human{"Danila", 20}
	DanilaAction := Action{Danila}
	Danila.Talk("Salut!")
	DanilaAction.Jump()
}
