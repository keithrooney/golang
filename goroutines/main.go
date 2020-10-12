package main

// A goroutine is a lightweight thread managed by the Go runtimes
// Goroutines run in the same address spacec, so access to shared memory must be synchronized
// Might be worth looking at the "sync" package

import (
	"fmt"
	"time"
)

// The evaluation of s happens in the current goroutine
// The execution of say happens in a new goroutine
func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("World")
	say("Hello")
}
