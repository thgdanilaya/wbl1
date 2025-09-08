package main

import "fmt"

func main() {
	var (
		s string
		n int
		v []string
	)
	m := make(map[string]bool)

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&s)
		v = append(v, s)
		m[s] = true
	}
	fmt.Println(v)
	for key := range m {
		fmt.Printf("%v ", key)
	}
}
