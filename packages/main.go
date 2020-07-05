// All Go programs are made up of packages.
package main

// The below imports are imports into a parenthesized, "factored" import statement
import (
	"fmt"
	"math/rand"
)

// The above is the same as the following ... 
// 
// import "fmt"
// import "math"
//
// It is good style to use the factored import statement.

func main() {
	fmt.Println("My favourite number is", rand.Intn(10))
}

