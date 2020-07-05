// Go has only one looping construct, the for loop.
package main

import "fmt"

func main() {

	// Unlike C or Java, there are parenthesis surrounding the three components.
        // The init statement
        // The condition statement
        // The post statement

	x := 0
	for i := 0; i < 10; i++ {
		x += i
	}
	fmt.Println(x)

	// The init and post statements are optional
	y := 1
	for ; y < 1000; {
		y += y
	}
	fmt.Println(y)

	// At this point we can drop the semi-colons in absence of an init and post statement
	z := 1
	for z < 1000 {
		z += z
	}
	fmt.Println(z)


	// Below is an example of an infinite loop
	// for {
	// }

}

