package main

var p *int

//2, d := test1()
// 1, test1()
func test1() *int {
	var a int
	return &a
}

// 2
func test2() {
	var b int

	p = &b
}

// 1
func test3() {
	var c int
	c += 1
}

func main() {
	d := test1()

	test2()

	test3()

	*d += 1
}
