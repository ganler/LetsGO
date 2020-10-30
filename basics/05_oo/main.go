package main

import (
	"fmt"
	"os"
	"time"
)

type worker struct {
	name string
	age  int
}

func hireWorker(name string, age int) worker {
	if age > 35 {
		fmt.Fprintf(os.Stderr, "Hi %s, We don't hire boys older than 35!", name)
		os.Exit(1)
	}
	return worker{name: name, age: age}
}

// Normal func: func [name](args...) ret_type{}
// Method func: func (self *type) [name](args...) ret_type{}
func (w *worker) Work() {
	fmt.Printf("%s is working under 996 mode.\n", w.name)
	time.Sleep(996 * time.Microsecond)
}

func main() {
	xzt := hireWorker("xzt", 23)
	xzt.Work()
	// You can also use `.` for pointers like:
	(&xzt).Work()

	mwish := hireWorker("fxw", 36)
	mwish.Work()
}
