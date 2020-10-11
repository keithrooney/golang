// An interface type is defined as a set of method signatures

package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

type I interface {
	M()
}

type T struct {
	S string
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
// Implicit interfaces decouplte the definition of an interface from its implementation,
// which could appear in any package without prearrangement.
func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func main() {

	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f

	// Vertex (the value type) doesn't implement Abser because Abs method is defined only on *Vertex (the pointer type)
	a = &v
	// a = v

	fmt.Println(a.Abs())

	// var i I = T{"hello"}
	// i.M()

	var i I
	// describe(i)
	// i.M()

	var t *T
	i = t
	describe(i)
	i.M()

	i = &T{"Hello"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()

	var x interface{}
	describeEmpty(x)

	x = 43
	describeEmpty(x)

	x = "hello"
	describeEmpty(x)

}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func describeEmpty(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
