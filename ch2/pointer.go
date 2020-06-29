package main

import "fmt"

func main() {
	a := 1
	b := 2

	swap(&a, &b)

	fmt.Printf("a %d, b %d\n", a, b)
}

func test() {
	// int *p = NULL;
	// p := nil 错误:use of untyped nil
	var p *int = nil

	// int a = 0;
	var a int = 0

	// p = &a;
	p = &a

	// *p = 1
	*p = 1

	// printf("%d", a);
	fmt.Println(a)

	// printf("%d", *p);
	fmt.Println(*p)
}

func swap(a *int, b *int) {
	tmp := *a
	*a = *b
	*b = tmp
}
