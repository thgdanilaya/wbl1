package main

import "fmt"

func main() {
	// ОБмен с помощью умножения, деления и xor
	a := 5
	b := 8
	fmt.Println(a, b)
	a = a * b
	b = a / b
	a = a / b
	fmt.Println(a, b)
	a = a ^ b
	b = b ^ a
	a = a ^ b
	fmt.Println(a, b)
}
