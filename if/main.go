package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	// If statements do not require rounded parentheses
	// However curly braces are required
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

// The if statement can start with a short statement to execute before the condition.
// This short statement is in the form of "v < lim" below.
// Variables declared by the statement are only in scope until the end of the if.
// Meaning, the usage of the variable v outside of the statement is invalid.
// Variables declared inside an if short statement is available inside any of the else blocks
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
	fmt.Println(pow(3, 2, 10), pow(3, 3, 20))
}

