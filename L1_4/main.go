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

func worker1(id int, wg *sync.WaitGroup, jobs <-chan job, ctx context.Context) { // Как я понимаю в таком случае у нас сразу прерывается обработка
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("worker %d terminated\n", id)
			return
		case job, ok := <-jobs:
			if !ok {
				fmt.Printf("worker %d empty and terminated\n", id)
				return
			}
			fmt.Printf("Worker id: %d %s\n", id, job.s)
		}
	}
}

func worker2(id int, wg *sync.WaitGroup, jobs <-chan job) { // А вот здесь уже очередь из сообщений до конца чиститься будет, так как в 76 строчке
	// при ctx.Done закроется канал с джобами, воркеры дочитают элементы из Jobs и стопнутся
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Worker id: %d %s\n", id, job.s)
	}
	fmt.Printf("worker %d empty and terminated\n", id)
}

func main() {
	var wg sync.WaitGroup
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
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
	jobs := make(chan job, 64) // добавил буффер
	in := make(chan string)

	go func() {
		defer close(in)
		sc := bufio.NewScanner(os.Stdin)
		for sc.Scan() {
			in <- sc.Text()
		}
	}()

	go func() {
		defer close(jobs)
		for {
			select {
			case <-ctx.Done():
				return
			case s, ok := <-in:
				if !ok {
					return
				}
				jobs <- job{s}
			}
		}
	}()

	for w := 0; w < workers; w++ {
		id := w + 1
		//go worker1(id, &wg, jobs, ctx)
		go worker2(id, &wg, jobs)
	}
	wg.Wait()
}
