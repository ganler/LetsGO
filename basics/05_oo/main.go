package main

import (
	"fmt"
	"math"
	"os"
	"time"
)

type worker struct {
	name string
	age  int
}

func hireWorker(name string, age int) worker {
	if age > 35 {
		fmt.Fprintf(os.Stderr, "[WARNING] Hi %s, We don't hire boys older than 35!", name)
	}
	return worker{name: name, age: age}
}

// Normal func: func [name](args...) ret_type{}
// Method func: func (self *type) [name](args...) ret_type{}
func (w *worker) Work() {
	fmt.Printf("%s is working under 996 mode.\n", w.name)
	time.Sleep(996 * time.Microsecond)
}

// About interface. (kinda like virtual functions in C++)
type shape interface {
	area() float64
}

type rect struct {
	w float64
	h float64
}

type circle struct {
	r float64
}

// Whether to use *: * for ref, non-* for copy.
// Everything in go is passing by value.
// But you can choose to pass a light pointer for reference.
func (r rect) area() float64 {
	return r.w * r.h
}

func (c circle) area() float64 {
	return math.Pi * c.r * c.r
}

func printArea(s shape) {
	fmt.Println(s, "area =", s.area())
}

func main() {
	xzt := hireWorker("xzt", 23)
	xzt.Work()
	// You can also use `.` for pointers like:
	(&xzt).Work()

	mwish := hireWorker("fxw", 36)
	mwish.Work()

	r := rect{w: 1, h: 2}
	c := circle{r: 1.5}
	printArea(r)
	printArea(c)
}
