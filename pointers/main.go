package main

import "fmt"

// A pointer holds the memory address of a value, like C++
// The type *T is a pointer to a T value, its zero value is nil
// 
// var p *int
//
// The & operator generates a pointer to its operand
// 
// i := 42
// p = &i
// 
// The * operator denotes the pointer's underlying value
// 
// fmt.Println(*p) 	// Read i through the pointer p, "dereferencing"
// *p = 21		// Set i through the pointer p, "indirecting"
//
// Notes
// 
// https://medium.com/@meeusdylan/when-to-use-pointers-in-go-44c15fe04eac
//
// * Passing pointers in Go is often slower than passing values
// * Go is a garbage collected language, much like that of Java.
// * Passing a pointer to a function means Escape Analysis has to be performed.
// 	** This is used to figure out if the variable should be stored on the heap or stack.
//	** You can check what the analysis is doing by running with 'go build -gcflags="-m"'
// 	** Using this flag, Go will tell you if a variable will escape to the heap or not
//	** If a variable does not escape to the heap, it lives on the stack.
//	** A stack does not require a garbage collector to clean up variables, it's just a push/pop
// * This analysis will add a bit of overhead but, in addition, the variable could be stored on the heap.
// * When you store a variable on the heap, you also lose time when GC is running.
// * Passing by value, you will always run on the stack, not incurring overhead in GC.
// * ... but when do we want to use pointers?
// 	** Copying large structs - benchmark for size of struct
//	** Mutability - pass by value will only mutate the passed value
//	** API consistency - if you use a pointer in place, you it elsewhere
//	** To signify true absence - default zero value of a pointer is nil signaliing true absence.
//
func main() {

	i, j := 42, 2701

	p := &i			// point to i
	fmt.Println(*p)		// read i through the pointer

	*p = 21			// set i through the pointer
	fmt.Println(i)		// see the new value of i

	p = &j			// point to j
	*p = *p / 37		// divide j through the pointer
	fmt.Println(j)

}

