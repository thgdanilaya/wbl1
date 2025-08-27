package main

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	// Самый очевидный и нулевой способ это просто дождаться, когда главная горутина закроется, тогда и все остальные тоже лягут отдыхать
	var wg sync.WaitGroup
	wg.Add(1)
	go func() { // Первый способ завершения горутины через ретурн
		defer wg.Done()
		for i := 0; i <= 2; i++ {
			fmt.Println("Hello World")
		}
		return
	}()

	// Второй способ через channel когда туда что-то придет
	wg.Wait()
	done := make(chan string)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-done:
				fmt.Println("Second Done")
				return
			default:
				fmt.Println("Wait for second")
				time.Sleep(1 * time.Second)
			}
		}
	}()
	time.Sleep(3 * time.Second)
	done <- "Hello World"
	defer close(done)

	// Третий способ через закрытие канала

	third := make(chan string)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			str, ok := <-third
			if !ok {
				fmt.Println("third")
				return
			}
			fmt.Println(str)
		}
	}()
	time.Sleep(3 * time.Second)
	close(third) // через 3 секунды закрываем канал и горутина тоже перестанет работать
	wg.Wait()
	// Четвертый способ завершения через контекст
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("context timeout")
				return
			default:
				fmt.Println("Wait for context")
				time.Sleep(1 * time.Second)
			}
		}
	}()
	wg.Wait()

	// С Мьютексом
	var mu sync.Mutex
	var stop bool
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			mu.Lock()
			s := stop
			mu.Unlock()
			if s {
				fmt.Println("stopped by mutex")
				return
			}
			fmt.Println("Wait for mutex")
			time.Sleep(1 * time.Second)
		}
	}()
	time.Sleep(3 * time.Second)
	mu.Lock()
	stop = true
	mu.Unlock()
	wg.Wait()

	// Atomic BOol

	var st atomic.Bool
	wg.Add(1)

	go func() {
		defer wg.Done()
		for {
			if st.Load() {
				fmt.Println("stopped by atomic")
				return
			}
			fmt.Println("Wait for atomic")
			time.Sleep(1 * time.Second)
		}
	}()
	time.Sleep(3 * time.Second)
	st.Store(true)
	wg.Wait()

}
