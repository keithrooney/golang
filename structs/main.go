package main

import "fmt"

// A struct is a collection of fields
type Vertex struct {
	X int
	Y int
}

func main() {
	// This is how you initialize a struct
	vertex := Vertex{1, 2}

	// This is how you access a struct's field
	vertex.X = 4

	fmt.Println(vertex.X)

	another := Vertex{4, 5}

	// struct fields can be accessed through a struct pointer.
	// Accessing a struct's field when we have the pointer can be written as (*p).X
	// However, Golang have shorten this to allow us to simply write p.X without dereferencing.
	something := &another
	something.X = 1e9

	fmt.Println(another)
	fmt.Println(something)

	v1 := Vertex{1, 2} // has type Vertex
	v2 := Vertex{X: 1} // Y:0 is implicit
	v3 := Vertex{}     // X:0 and Y:0 
	p := &Vertex{1, 2}  // has type *Vertex or pointer to Vertex

	fmt.Println(v1, v2, v3, p)

}

