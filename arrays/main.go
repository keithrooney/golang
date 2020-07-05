package main

import "fmt"

func main() {

	// Declares an array of 2 strings
	// An array's length is part of its type so they cannot be resized
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"

	fmt.Println(a[0], a[1])
	fmt.Println(a)

	// Arrays also support the short assignment
	primes := [6]int{2, 3, 5, 7, 11, 13}

	fmt.Println(primes)
}

