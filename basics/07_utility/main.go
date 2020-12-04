package main

import (
	"fmt"
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
}
