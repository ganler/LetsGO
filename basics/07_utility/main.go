package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// timer as a notifier. Do sth only when timer is ok!
	timer := time.NewTimer(1 * time.Second)
	go func() {
		<-timer.C
		fmt.Println("It is time now!")
	}()
	fmt.Println("Timer stop successful? :=> ", timer.Stop()) // timer is stopped!
	time.Sleep(1 * time.Second)                              // last goroutine is stuck...

	// tickers are for when you want to do something repeatedly at regular intervals.
	ticker := time.NewTicker(500 * time.Millisecond)
	doneFlag := make(chan bool)
	go func() {
		for {
			select {
			case <-doneFlag:
				return
			case t := <-ticker.C:
				fmt.Println("Tick @ ", t)
			}
		}
	}()
	time.Sleep(2 * time.Second)
	ticker.Stop()
	doneFlag <- true

	// The defer key word:
	func() {
		for i := 0; i < 3; i++ {
			defer fmt.Println(i) // Defer is like a stack in current function. FILO.
		}
	}()

	// Wait Group: Kinda like fork and join.
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1) // Add one more task; (it is like an atomic counter)
		go func(id int, wg *sync.WaitGroup) {
			defer wg.Done()
			fmt.Println("Wait Group Task ID =", id)
		}(i, &wg)
	}
	wg.Wait()
}
