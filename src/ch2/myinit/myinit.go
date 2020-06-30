package main

import "fmt"

var test [128]int

func init() {
	for i := range test {
		test[i] = i
	}
}

func main() {
	for i := range test {
		fmt.Println(test[i])
	}
}
