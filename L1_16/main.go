package main

import "fmt"

func quickSort(a []int) []int {
	if len(a) < 2 {
		return a
	}
	pivot := a[len(a)/2]
	var less []int
	var greater []int
	var equal []int
	for _, value := range a {
		if value < pivot {
			less = append(less, value)
		} else if value > pivot {
			greater = append(greater, value)
		} else {
			equal = append(equal, value)
		}
	}
	return append(append(quickSort(less), equal...), quickSort(greater)...)
}

func main() {
	var n, a int
	var numbers []int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a)
		numbers = append(numbers, a)
	}

	result := quickSort(numbers)
	fmt.Println(result)
}
