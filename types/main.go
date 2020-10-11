// Go basic types are ...
//
// bool
// string
// int, int8, int16, int32, int64
// uint, uint8, uint16, uint32, uint64, uintptr
// byte (alias for uint8)
// rune (alias for int32, represents a Unicode code point)
// float32, float64
// complex64, complex128
//
// int, uint and uintptr types are usually 32 bits wide on 32 bit systems and 64 bits wide on 64-bit systems.
// When you need an integer value you should use int, unless you have a specific reason to use a sized or unsigned integer type

package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	// Variables without an explicit initial value are given zero values
	// 0 for numeric types
	// false for boolean types
	// "" empty string for strings
	var i int
	var f float64
	var b bool
	var s string

	fmt.Printf("%v %v %v %q\n", i, f, b, s)

	// Unlike C, in Go, items of different type require explicit conversions
	var x, y int = 3, 4
	var float float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(float)

	fmt.Println(x, y, z)

	// When the right hand side of the declaration is typed, the new variable is of that same type
	var l int
	m := i

	fmt.Printf("Type: %T\n", l)
	fmt.Printf("Type: %T\n", m)

	// ... but when the right hand side contains an utyped numeric constant,
	// the variable maybe an int, float64 or complex128, depending on precision of constant
	p := 42
	q := 3.142
	r := 0.867 + 0.5i

	fmt.Printf("Type: %T\n", p)
	fmt.Printf("Type: %T\n", q)
	fmt.Printf("Type: %T\n", r)

	var one interface{} = "hello"

	// Below is a type assertion.
	// It asserts the interface value holds the concrete type "string"
	// It assigns the underlying "string" value to the variable t
	// If m is not a string type, the statement will panic
	two := one.(string)
	fmt.Println(two)

	two, ok := one.(string)
	fmt.Println(two, ok)

	three, ok := one.(float64)
	fmt.Println(three, ok)

	// three = one.(float64) // this will panic
	// fmt.Println(three)

}
