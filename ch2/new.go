package main

// int *GetInt1() {
//     int *p = new int;
//     return p;
// }
//
// int *GetInt2() {
//     int a;
//     return &a;
// }

func GetInt1() *int {
	var p *int = new(int)
	return p
}

func GetInt2() *int {
	var a int
	return &a
}

func main() {
	// int *p = new int;

	var p *int = new(int)
}
