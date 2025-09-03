package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n)
	fmt.Scan(&m)
	var a []int
	var b []int
	var val int
	intersect := make(map[int]bool)

	for i := 0; i < n; i++ {
		fmt.Scan(&val)
		a = append(a, val)
	}
	for i := 0; i < m; i++ {
		fmt.Scan(&val)
		b = append(b, val)
	}

	for _, i := range a {
		intersect[i] = false
	}
	for _, i := range b {
		if _, ok := intersect[i]; ok {
			intersect[i] = true
		}
	}
	for key, value := range intersect {
		if value == true {
			fmt.Println(key)
		}

	}
}
