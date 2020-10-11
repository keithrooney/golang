package main

import (
	"fmt"
	"runtime"
	"time"
)

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about this type %T!\n", v)
	}
}

func main() {

	fmt.Print("Go runs on ")
	// Unlike Java, Go only runs the selected case, not all the cases that follow.
	// This means the break statement is not actually required.
	// Another difference is switch cases need not be constants, and the values involved need not be integers.
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}

	// Switch evaluates cases from top to bottom, stopping when a case succeeds
	// 2009-11-10 23:00:00 UTC is the date and time of Golangs birthday.
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today!")
	case today + 1:
		fmt.Println("Tomorrow!")
	case today + 2:
		fmt.Println("In two days!")
	default:
		fmt.Println("In a few days!")
	}

	// Switch can be written without a condition, it is the same as switch true
	// This way can be a clean way to write a long if-then-else chains.
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good evening!")
	}

	do(21)
	do("hello")
	do(true)

}
