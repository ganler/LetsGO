package main

import (
	"fmt"
	"math"
)

func main() {
	// Comments A
	/* Comments B */

	// Values: https://gobyexample.com/values
	fmt.Println("go" + "lang")
	fmt.Println("1+1 =", 1+1)
	fmt.Println("true =", true)
	// This part is quite the same as C++/Java.

	// Variables: https://gobyexample.com/variables
	var a = "var a"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	d := "d := xxx"
	fmt.Println(d)

	// Variables declared without a corresponding initialization are zero-valued.
	// For example, the zero value for an int is 0.
	var uninit int
	fmt.Println("uninit [int] var =", uninit)

	// Constant value.
	const n = 10
	fmt.Println("10 / 3 =", n/3)
	fmt.Println("10 / 3. =", n/3.)

	// A numeric constant has no type until it’s given one, such as by an explicit conversion.
	// e.g., `n` can be of any precision.
	// but `const n int = 10` means that n is of the `int` type.
	const f = 10312310 / 10.
	// Warning: You cannot truncate a non-int float var into int var.
	// Or you may get: .\main.go:39:39: constant truncated to integer
	fmt.Println("int64(", f, ") =", int64(f))
	fmt.Println("SIN(n)", math.Sin(n))

	// Convert Numeric Types.
	fmt.Println("round(0.51) =", math.Round(0.51))
	/*
		Round(±0) = ±0
		Round(±Inf) = ±Inf
		Round(NaN) = NaN
	*/
	// !: You cannot convert a const float into int numbers.
	// You must turn it into var first.
	// See: https://golang.org/ref/spec#Conversions
	const cv = 1.9
	v := cv
	fmt.Println("int(1.9) =", int(v)) //  1
}
