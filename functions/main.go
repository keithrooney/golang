package main

import (
	"fmt"
	"math"
)

func add(x int, y int) int {
	return x + y
}

// When two or more consecutive named function parameters share a type,
// you can omit the type from all but the last, meaning our function becomes ...
// func add(x, y int) ....

// A function can return any number of results, as seen below.
func swap(x, y string) (string, string) {
	return y, x
}

// Go's return values can be named, treated as variables defined at the top of the function.
// These names document the meaning of the return values.
// We can also omit a return argument in this case, it will return the named return values.
// The above is called "naked" return.
// These naked returns should only be used in short functions, harming readibility in longer functions.
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

// Function values, simply put, a function can be a value passed to a function
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

// An example of a closure in Go
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {

	fmt.Println(add(42, 33))

	a, b := swap("hello", "world")
	fmt.Println(a, b)

	fmt.Println(split(17))

	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))

	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-1*i),
		)
	}

}
