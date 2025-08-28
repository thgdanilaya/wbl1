package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	fmt.Println("_____Output map with RWMutex_____")
	casualMap()
	fmt.Println("_____Output map with syncMap_____")
	syncMap()
}

func casualMap() {
	m := make(map[string]int)
	var rw sync.RWMutex
	var wg sync.WaitGroup

	const writers = 2
	const writes = 50
	const readers = 2
	const reads = 50

	for w := 0; w < writers; w++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for i := 0; i < writes; i++ {
				key := fmt.Sprintf("key-w%d-i%d", id, i)
				rw.Lock() // Локаем для доступа ток для этой горутины
				m[key] = i
				rw.Unlock()
			}
		}(w)
	}

	for r := 0; r < readers; r++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < reads; i++ {
				key := fmt.Sprintf("key-w%d-i%d", rand.Intn(writers), rand.Intn(writes))
				rw.RLock()
				v, ok := m[key] // Будем просто читать без вывода
				if ok {
					fmt.Println(v, key)
				}
				rw.RUnlock()
			}
		}()
	}
	wg.Wait()
}

func syncMap() {
	var sm sync.Map
	var wg sync.WaitGroup

	const writers = 2
	const writes = 50
	const readers = 2
	const reads = 50

	for w := 0; w < writers; w++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for i := 0; i < writes; i++ {
				key := fmt.Sprintf("key-w%d-i%d", id, i)
				sm.Store(key, i)
			}
		}(w)
	}

	for r := 0; r < readers; r++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < reads; i++ {
				key := fmt.Sprintf("key-w%d-i%d", rand.Intn(writers), rand.Intn(writes))
				v, ok := sm.Load(key)
				if ok {
					fmt.Println(key, v)
				}
			}
		}()
	}
	wg.Wait()
}
