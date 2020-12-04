package main

import (
	"fmt"
	"time"
)

func strChange(str string) string {
	return "[" + str + "]"
}

func sumChan(s []int, ch chan int) {
	ret := 0
	for _, v := range s {
		ret += v
	}
	ch <- ret
}

func fibonacci(datachan chan<- int /* chan<- push-only */, quit <-chan int /* <-chan pop-only */) {
	x, y := 0, 1
	fmt.Println("begin fibonacci...")
	for {
		select {
		// Select the upcoming case. If there're multiple ready ones, do it at random.
		case datachan <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit...")
			return
		default: // Runs if no one is ready...
			fmt.Println("No case is ready...")
			time.Sleep(50 * time.Millisecond)
		}
	}
}

func timeoutMessage(flag <-chan bool) {
	select {
	case <-flag:
		fmt.Println("NOT TIME OUT!")
	case <-time.After(3 * time.Second):
		fmt.Println("TIMEOUT -> 3s!")
	}
}

func main() {
	// A goroutine is a lightweight thread managed by the Go runtime.
	// The parameters are evaluated in the current thread, but the will
	// function will be executed in the GOthread.
	go fmt.Println(strChange("world"))
	fmt.Println(strChange("hello"))

	// Channels
	// Channels are a typed conduit through which you can send and receive
	// values with the channel operator, <-.
	// ch <- v
	// v := <- ch
	ch := make(chan int)
	s := []int{1, 2, 3, 4, 5, 6}
	// The output order is not fixed.
	// By default, the buffer size of channel is one that it will block sending util
	// it is empty again.
	go sumChan(s[:len(s)/2], ch)
	go sumChan(s[len(s)/2:], ch)
	x, y := <-ch, <-ch // receive values.
	// By default, sends and receives block until the other side is ready.
	// This allows goroutines to synchronize without explicit locks or condition variables.

	fmt.Println(x, y, x+y)

	// Buffered Channels
	bch := make(chan int, 2)
	bch <- 1
	bch <- 2
	fmt.Println(<-bch, <-bch)

	// Range and Close
	// A sender can close a channel to indicate that no more values will be sent.
	go func() {
		bch <- 1
		bch <- 2
		// Receivers can test whether a channel has been closed by assigning a second
		// parameter to the receive expression: after
		close(bch)
		// Only the sender should close a channel, never the receiver. Sending on a
		// closed channel will cause a panic.
		// Channels aren't like files;
		// you don't usually need to close them.
		// Closing is only necessary when the receiver must be told there are no more
		// values coming, such as to terminate a range loop.
	}()

	for i := range bch {
		fmt.Println(i)
	}

	// Or if you want to check it manually.
	// v, ok := <- bch
	datachan := make(chan int)
	quit := make(chan int)
	go func() {
		for range [5]int{} {
			fmt.Println(<-datachan)
		}
		quit <- 1
	}()
	fibonacci(datachan, quit)

	// Use channel as a tool for synchronization;
	doneFlag := make(chan bool)
	go func(flag chan<- bool) { // chan<- means this channel is `push-only`;
		// Do sth...
		fmt.Println("[BeforeSync] Do sth before set the sync flag;")
		flag <- true
	}(doneFlag)
	<-doneFlag // Sync.
	fmt.Println("[AfterSync] This happens after sync!")

	// timeout flag;
	reqFlag := make(chan bool)
	go timeoutMessage(reqFlag)
	time.Sleep(2 * time.Second)
	go timeoutMessage(reqFlag)
	time.Sleep(2 * time.Second)
	reqFlag <- true
	close(reqFlag)
}
