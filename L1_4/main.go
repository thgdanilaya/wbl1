/*
Был выбран для завершения всех горутин-воркеров Context, так как его легко провести для воркеров.
Ну и я еще вычитал, что можно отменять контекст не только по сигналу, но и по таймеру, к примеру.
*/

package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
)

type job struct {
	s string
}

func worker(id int, wg *sync.WaitGroup, jobs <-chan job, ctx context.Context) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("worker %d terminated\n", id)
			return
		case job, ok := <-jobs:
			if !ok {
				fmt.Printf("worker %d terminated\n", id)
				return
			}
			fmt.Printf("Worker id: %d %s\n", id, job.s)
		}
	}
}

func main() {
	var wg sync.WaitGroup
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

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
		go worker(id, &wg, jobs, ctx)
	}
	wg.Wait()
}
