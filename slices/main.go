package main

import (
	"fmt"
	"strings"
)

// Slices are much more common than arrays
func main() {

	// Below is an array of integers with fixed length 6
	primes := [6]int{2, 3, 5, 7, 11, 13}

	// A slice is a dynamically-sized flexible view into the elements of an array
	// A slice is formed by specifying a low and high bound, essentially a range of values from an array
	var s []int = primes[1:4]

	fmt.Println(s)

	// A slice does not store any data, it describes a section of the underlying data
	// Modifying the slice will in turn modify the underlying data structure
	s[0] = 17
	s[1] = 19

	fmt.Println(s)
	fmt.Println(primes)

	// A slice literal is like an array literal without the length
	p := [6]int{2, 3, 5, 7, 11, 13}
	q := []int{2, 3, 5, 7, 11, 13}

	fmt.Println(p)
	fmt.Println(q)

	m := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(m)

	// When slicing, you may omit the high or low bounds to use their defaults instead
	a := []int{2, 3, 5, 7, 11, 13}

	a = a[1:4]
	fmt.Println(a)

	// The default is zero for the lowwer bound ...
	a = a[:2]
	fmt.Println(a)

	// ... and the higher bound default is the length of the slice
	a = a[1:]
	fmt.Println(a)

	// A slice has both a length and a capacity
	// The length of a slice is the number of elements it contains
	// The capacity of a slice is the number of elements in the underlying array, counting from the first element.
	b := []int{1, 2, 3, 4, 5, 6}

	fmt.Printf("len=%d cap=%d %v\n", len(b), cap(b), b)

	b = b[:0]
	fmt.Printf("len=%d cap=%d %v\n", len(b), cap(b), b)

	// You can extend a slice's length by re-slicing it, providing it has sufficient capacity.
	// Remeber, the capacity of a slice is the number of elements in the underlying array
	b = b[:4]
	fmt.Printf("len=%d cap=%d %v\n", len(b), cap(b), b)

	// The zero value of a slice is nil
	var c []int
	fmt.Println(c, len(c), cap(c))
	if c == nil {
		fmt.Println("nil!")
	}

	x := make([]int, 5)
	printSlice("x", x)

	y := make([]int, 0, 5)
	printSlice("y", y)

	z := b[:2]
	printSlice("z", z)

	w := z[2:5]
	printSlice("w", w)

	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	var slice []int
	printSlice1(slice)

	slice = append(slice, 0)
	printSlice1(slice)

	slice = append(slice, 1)
	printSlice1(slice)

	slice = append(slice, 2, 3, 4)
	printSlice1(slice)

}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}

func printSlice1(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
