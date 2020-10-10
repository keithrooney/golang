package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	counts := make(map[string]int)
	words := strings.Split(s, " ")
	for index := range words {
		word := words[index]
		value, _ := counts[word]
		counts[word] = value + 1
	}
	return counts
}

func main() {
	fmt.Println(WordCount("Hello Keith, today is your day!"))
}
