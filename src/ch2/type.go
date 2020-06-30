package main

import "fmt"

// typedef signed int int32;

type Date int32
type Time int32

func main() {
	var currentDate Date
	var currentTime Time

	currentDate = 29
	currentTime = 24

	fmt.Printf("Date %d, Time %d\n", currentDate, currentTime)
}
