package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"
)

type job struct {
	s string
}

func main() {
	var wg sync.WaitGroup

	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: app <workers>")
		os.Exit(1)
	}

	workers, err := strconv.Atoi(os.Args[1])
	if err != nil || workers <= 0 {
		fmt.Fprintln(os.Stderr, "workers must be positive integer")
		os.Exit(1)
	}

	wg.Add(workers)
	jobs := make(chan job)
	in := make(chan string)

	go func() {
		defer close(in)
		sc := bufio.NewScanner(os.Stdin)
		for sc.Scan() {
			text := sc.Text()
			in <- text
		}
	}()

	go func() {
		defer close(jobs)
		for s := range in {
			jobs <- job{s}
		}
	}()

	for w := 0; w < workers; w++ {
		id := w + 1
		go func(id int) {
			defer wg.Done()
			for j := range jobs {
				fmt.Printf("Worker id: %d %s\n", id, j.s)
			}
		}(id)
	}
	wg.Wait()
}
