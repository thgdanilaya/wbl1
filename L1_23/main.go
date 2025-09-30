package main

import "fmt"

func main() {
	var n, idx int
	fmt.Scan(&n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	fmt.Scan(&idx)
	if idx < 0 || idx >= len(a) {
		fmt.Println("Index out of range")
		return
	}
	copy(a[idx:], a[idx+1:])
	a[len(a)-1] = 0 // на всякий случай затрем нафиг что тут будет лежать
	a = a[:len(a)-1]
	fmt.Println(a)
}
