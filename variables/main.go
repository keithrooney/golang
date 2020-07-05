package main

import "fmt"

// "var" is used to declare variables.
// Variables can be declared as a list
// The type is the last statement in the list
// var c, python, java bool

// A var can include initializers, like below.
// If an initializer is present, the type can be omitted.
var x, y int = 1, 2

// var c, python, java = true, false, "no!"

func main() {
	var i, j int = 1, 2

	// The short assignment := can be used in place of a var declaration with implicit type
	// Outside a function, every statement beings with a keyword, i.e. var, func and so on
	// Therefore the short assginment operator is not available outside
	k := 3

	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)

}

