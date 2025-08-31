package main

import "fmt"

func main() {
	var n int64
	var i int64
	var val int
	fmt.Println("Enter n: ")
	fmt.Scan(&n)
	fmt.Println("Enter i: ")
	fmt.Scan(&i)
	fmt.Println("Enter val(only 0 or 1): ")
	fmt.Scan(&val)
	mask := int64(1) << i
	if val == 1 {
		n |= mask
	} else {
		n &^= mask
	}
	fmt.Printf("%b\n", n)
}
