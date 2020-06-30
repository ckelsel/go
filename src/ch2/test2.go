package main

import "os"
import "fmt"

var dir string

func main() {
	{
		dir, err := os.Getwd()
		if err != nil {
			return
		}

		fmt.Println(dir)
	}
	fmt.Println(dir)
}
