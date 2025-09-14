package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	reverted := []rune(s)

	for i, j := 0, len(reverted)-1; i < j; i, j = i+1, j-1 {
		reverted[i], reverted[j] = reverted[j], reverted[i]
	}
	fmt.Println(string(reverted))
}
