package main

import "fmt"

// func (arg0 T0, arg1 T1, ...) RetType {}
func add(a int, b int) int {
	return a + b
}

// https://gobyexample.com/functions
// When you have multiple consecutive parameters of the same type,
// you may omit the type name for the like-typed parameters up to the final parameter that declares the type.
func add3(a, b, c int) int {
	return a + b + c
}

func sum(ns ...int) int {
	total := 0
	for _, n := range ns {
		total += n
	}
	return total
}

// Closures
func intSeq() func() int {
	i := 0
	return func() int {
		i++ // i is by this function. i.e., this functor owns a unique copy of i.
		return i
	}
}

/* If you are a C++ programmer...
auto intSeq() {
    int i = 0;
    return [=]() mutable {
        return ++i;
    };
}
*/

func main() {
	fmt.Println(add(1, 2))
	fmt.Println(add3(1, 2, 3))

	// Variadic functions can be called in the usual way with individual arguments.
	fmt.Println(sum(1, 2, 3, 4, 5, 6))
	// If you already have multiple args in a slice,
	// apply them to a variadic function using func(slice...) like this.
	ints := []int{2, 2, 2}
	fmt.Println(sum(ints...))

	// Closures
	nextint := intSeq()
	// The captured value is unique in each functor.
	fmt.Println(nextint(), nextint(), nextint())
	fmt.Println(intSeq()()) // This is a new functor starting from i = 0.
}
