package main

// A name is exported if it begins with a capital letter.
// For example, Pizza is an exported name, as is Pi, from the math package.
// pizza and pi do not start with a capital, so they're not exported.
import (
	"fmt"
	"math"
)

// You cannot refer to names which are not exported, for example, pizza or pi.
// The below is an example of a failure, running this will result in failure.
// func main() {
// 	fmt.Println(math.pi)
// }

// Below is the functioning example of using a valid export.
func main() {
	fmt.Println(math.Pi)
}

