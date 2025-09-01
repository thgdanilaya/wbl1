package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	var wg sync.WaitGroup
	m := make(map[int64][]float64)
	tempChannel := make(chan float64)
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(tempChannel)
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			select {
			case <-ctx.Done():
				log.Println("interrupt")
				return
			default:
			}
			temperature, err := strconv.ParseFloat(scanner.Text(), 64)
			if err != nil {
				log.Printf("skip %q: ", err)
				continue
			}
			select {
			case <-ctx.Done():
				return
			case tempChannel <- temperature:
			}
		}
		if err := scanner.Err(); err != nil {
			log.Println("scanner.Err:", err)
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for temperature := range tempChannel {
			key := int64(temperature/10) * 10
			m[key] = append(m[key], temperature)
		}
	}()
	wg.Wait()
	for key, value := range m {
		fmt.Println(key, value)
	}
}
