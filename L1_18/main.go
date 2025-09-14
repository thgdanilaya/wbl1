package main

import (
	"sync"
	"sync/atomic"
)

type count struct {
	c atomic.Int64
}

type count2 struct {
	mu sync.Mutex
	c  int64
}

func (c *count2) Inc() {
	c.mu.Lock()
	c.c++
	c.mu.Unlock()
}

func main() {
	var wg sync.WaitGroup
	var c count
	var c2 count2
	workers := 3

	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				c.c.Add(1)
			}
		}()
	}
	wg.Wait()
	println(c.c.Load())

	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				c2.Inc()
			}
		}()
	}
	wg.Wait()
	println(c2.c)

}
