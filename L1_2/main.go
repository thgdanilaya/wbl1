package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	sc := bufio.NewScanner(os.Stdin)

	var Numbers []int
	for {
		if !sc.Scan() {
			break
		}
		line := strings.TrimSpace(sc.Text())
		if line == "" {
			break
		}
		for _, f := range strings.Fields(line) {
			if n, err := strconv.Atoi(f); err == nil {
				Numbers = append(Numbers, n)
			} else {
				continue
			}
		}
	}

	wg.Add(len(Numbers))
	for i := range Numbers {
		i := i
		go func() {
			defer wg.Done()
			Numbers[i] = Numbers[i] * Numbers[i]
		}()
	}
	wg.Wait()

	fmt.Println(Numbers)
}
