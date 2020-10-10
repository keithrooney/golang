package main

import "fmt"

func fibonacci() func() int {
	first, second := 0, 1
	return func() int {
		ret := second
		second, first = first, ret+first
		return ret
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
