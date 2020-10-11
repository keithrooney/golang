package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Removal of the * (pointer) reference means it is a value receiver.
// A value receiver means the method would operate on a copy of the original Vertex value
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {

	v0 := Vertex{3, 4}
	fmt.Println(v0.Abs())
	fmt.Println(Abs(v0))

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())

	v1 := Vertex{3, 4}
	fmt.Println(v1.Abs())
	// Go interprets the statement v.Scale(5) as (&v).Scale(5) since Scale method has a pointer receiver
	v1.Scale(10)
	fmt.Println(v1.Abs())
	eg := &v1
	fmt.Println(eg.Abs())

	v2 := Vertex{3, 4}
	fmt.Println(Abs(v2))
	Scale(&v2, 10) // & means pass the reference
	fmt.Println(Abs(v2))

	// Functions that take value arguments must take a value of that specific type
	v3 := Vertex{5, 5}
	fmt.Println(v3.Abs())
	fmt.Println(Abs(v3))

	// Methods with value receiveres take either a value or a pointer as the receiver when called
	v4 := &Vertex{5, 5}
	fmt.Println(v4.Abs())
	fmt.Println(Abs(*v4))

	// There are two reasons for pointer receivers, as noted in the Golang docs :-
	// 1. The method can modify the value its receiver points to
	// 2. To avoid copying the value on each method, which is efficient if the receiver is a large struct
	// Generally, all methods on a given type should have either a value or pointer receiver, but not a mixture

}
