package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	m := make(map[rune]bool)
	fmt.Scan(&s)
	s = strings.ToLower(s)

	for _, v := range s {
		if !m[v] {
			m[v] = true
		} else {
			fmt.Println("False")
			return
		}
	}
	fmt.Println("True")

}
