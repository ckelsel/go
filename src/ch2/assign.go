package main

import "os"
import "fmt"

func main() {
	// var str string
	// var value int
	//
	// str = "abc"
	// value = 1

	var x, y int
	x, y = 0, 1
	// x = 0
	// y = 1

	// printf
	fmt.Printf("x %d, y %d\n", x, y)

	x, y = y, x
	// z := x
	// x = y
	// y = z

	fmt.Printf("x %d, y %d\n", x, y)

	file, err := os.Open("1.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	file.Close()

}
