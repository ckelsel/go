package main

import "fmt"

// 包级别的变量
var packageVar int

func main() {
	// 函数级别的变量
	var funcVar int

	// i for语法块级别的变量
	for i := 0; i < 10; i++ {
		var funcVar int

		packageVar = i

		funcVar = i

		fmt.Printf("i = %d\n", i)
	}

	fmt.Printf("funcVar = %d\n", funcVar)

	fmt.Printf("packageVar = %d\n", packageVar)
}
