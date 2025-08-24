package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	ttl, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("TTL must be an integer")
		os.Exit(1)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(ttl)*time.Second)
	defer cancel()

	channel := make(chan string) // Здесь не будем добавлять в WaitGroup, потому что по истечении таймера консоль все еще будет ожидать ввода.
	go func() {
		defer wg.Done()
		defer close(channel)
		sc := bufio.NewScanner(os.Stdin)
		for sc.Scan() {
			select {
			case <-ctx.Done():
				return
			case channel <- sc.Text():
			}
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("timeout")
				return
			case str, ok := <-channel:
				if !ok {
					return
				}
				fmt.Println(str)
			}
		}
	}()
	wg.Wait()
}
