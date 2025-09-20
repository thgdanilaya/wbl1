package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func reverse(runes []rune, start, end int) {
	for start < end {
		runes[start], runes[end] = runes[end], runes[start]
		start++
		end--
	}
}

func main() {
	scan := bufio.NewReader(os.Stdin)
	line, err := scan.ReadString('\n')
	if err != nil {
		return
	}
	runes := []rune(strings.TrimSpace(line))
	reverse(runes, 0, len(runes)-1)
	start := 0
	for i := 0; i <= len(runes); i++ {
		if i == len(runes) || runes[i] == ' ' {
			reverse(runes, start, i-1)
			start = i + 1
		}
	}
	fmt.Println(string(runes))
}
