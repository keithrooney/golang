package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	for i := int(x / 2); i >= 0; i-- {
		if number := i << 1; number <= int(x) {
			return float64(i)
		}
	}
	return -1
}

func main() {
	// Fails for the negative numbers
	// fmt.Println(Sqrt(-137))
	// fmt.Println(Sqrt(-3))
	// fmt.Println(Sqrt(-2))
	fmt.Println(Sqrt(0))
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(3))
	fmt.Println(Sqrt(8))
	fmt.Println(Sqrt(137))
}
