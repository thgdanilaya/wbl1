package main

import (
	"fmt"
)

func binSearch(a []int, research int) int {
	low := 0
	high := len(a)
	for low < high {
		mid := (low + high) / 2
		if a[mid] > research {
			high = mid
		} else if a[mid] < research {
			low = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

func main() {
	var (
		n        int
		research int
	)
	fmt.Scan(&n)
	numbers := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&numbers[i])
	}
	fmt.Scan(&research)
	res := binSearch(numbers, research)
	fmt.Println(res)
}
