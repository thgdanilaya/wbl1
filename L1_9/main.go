package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

func generateRandomNumber() int {
	return rand.Intn(100)
}

func main() {
	log.Println("Start")
	var wg sync.WaitGroup
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	first := make(chan int)
	second := make(chan int)
	wg.Add(3)
	go func() {
		defer wg.Done()
		defer close(first)
		for {
			select {
			case <-ctx.Done():
				log.Println("timed out")
				return
			case first <- generateRandomNumber():
				time.Sleep(2 * time.Second)
			}

		}
	}()
	go func() {
		defer wg.Done()
		defer close(second)
		for x := range first {
			second <- x * 2
		}
	}()
	go func() {
		defer wg.Done()
		for y := range second {
			fmt.Printf("%d\n", y)
		}
	}()
	wg.Wait()
	log.Println("done")
}
