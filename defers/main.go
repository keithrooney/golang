package main

import "fmt"

func main() {

	// A defer will defer the execution of a function until the surrounding function returns.
	// The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.
	defer fmt.Println("World!")
	fmt.Print("Hello ")

	// You can stack defers, they can be pushed onto a stack.
	// The stacked defers are executed in last-in first-out order.
	fmt.Println("Counting ... ")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("Done!")

}

