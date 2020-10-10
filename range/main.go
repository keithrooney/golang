package main

import "fmt"

func main() {

	for i, v := range []int{1, 2, 4, 8, 16, 32, 64, 128} {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	pow := make([]int, 10)
	for i := range pow {
		// fmt.Println(uint(i))
		pow[i] = 1 << uint(i) // == 2**i
	}

	fmt.Println(pow)

	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}

}
