package main

import (
	"fmt"
	"sync"
	"time"
)

func sleep(duration time.Duration) {
	deadline := time.Now().Add(duration)
	for time.Until(deadline) > 0 {
	}
}

func timerSleep(duration time.Duration) {
	timer := time.NewTimer(duration)
	<-timer.C
}

func main() {
	var d int
	fmt.Scan(&d)
	wg := sync.WaitGroup{}
	sec := time.Duration(d) * time.Second
	wg.Add(1)
	wg.Wait()
	go func() {
		defer wg.Done()
		fmt.Println("ждем функцию sleep")
		sleep(sec)
		fmt.Println("Дождались, теперь ждем timerSleep")
		timerSleep(sec)
		fmt.Println("voilà, monsieur, tout a fini")
	}()
}
