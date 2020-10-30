package main

import (
	"fmt"
	"time"
)

func main() {
	// For loop: initial/condition/after
	for i := 0; i < 10; i++ {
		fmt.Print(i)
	}
	fmt.Print("\n")

	// While True
	for {
		break
	}

	// While
	i := 0
	for i < 10 {
		fmt.Print(i)
		i++
	}
	fmt.Print("\n")

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	whatAmI := func(i interface{}) {
		fmt.Print("For var: ", i, " | ")
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Other type: %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")

	// Array
	var arr [5]int
	fmt.Println("empty arr: ", arr)
	arr = [5]int{1, 2, 3, 4, 5} // Note that := is used for defining new variables.
	fmt.Println("filled arr: ", arr)

	// Slices
	slice := make([]int, 3)

	// Similar to arrays: Indexing
	slice[0] = 1
	slice[1] = 2
	slice[2] = 3
	fmt.Println("SLICE:", slice, "Length =", len(slice))

	// Only for slice: Append function
	slice = append(slice, 4)
	fmt.Println("SLICE:", slice, "Length =", len(slice))

	slice = append(slice, 5, 6)
	fmt.Println("SLICE:", slice, "Length =", len(slice))

	// Slice syntax
	fmt.Println("slice[:3]", slice[:3]) // Just like python.
	fmt.Println("slice[1:]", slice[1:])

	// Concise slice definition
	t := []int{1, 2, 3}
	fmt.Println(t)

	// Pass by reference
	tref := t
	tref[0] = 666
	fmt.Println(t, tref)

	// Pass by value
	tcopy := make([]int, len(t))
	copy(tcopy, t) // copy(dst, src)
	tcopy[0] = 0
	fmt.Println(t, tcopy)

	// 2D slice
	ss := make([][]int, 3) // [[], [], []] The zero-val of slice : []
	fmt.Println(ss)
	for i := 0; i < len(ss); i++ {
		ss[i] = make([]int, i+1)
		for j := 0; j <= i; j++ {
			ss[i][j] = i + j
		}
	}

	fmt.Println(ss)

	// Maps
	m := make(map[string]int)
	// Or simply: m := map[string]int{"k0": 0, "k1": 1}
	m["k0"] = 0
	m["k1"] = 1

	fmt.Println(m, "with length =", len(m))

	delete(m, "k0")
	fmt.Println("After delete,", m, "with length =", len(m))

	// Ranges
	arr4 := [4]int{1, 2, 3, 4}
	for i, num := range arr4[:3] {
		fmt.Println("index:", i, " | value:", num)
	}

	for k, v := range map[string]string{"a": "apple", "b": "banana"} {
		fmt.Printf("%s -> %s\n", k, v)
	}
}
