package main

import "fmt"

// The below is a constant, which can be a string, boolean or numeric value
// Constants cannot be declared using the short assignment, or :=
const Pi = 3.14

// An untyped constant takes the type needed by it's context.
const (
	Big = 1 << 100
	Small = Big >> 99
)

// An int can store a maximum of 64-bit integer, and sometimes less.
func needInt(x int) int {
	return x * 10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {

	const world = "World!"
	fmt.Println("Hello,", world)
	fmt.Println("Happy", Pi, "Day")

	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
	// fmt.Println(needInt(Big))

}

