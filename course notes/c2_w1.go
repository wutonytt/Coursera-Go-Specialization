package main

import "fmt"

func PrintHello() {
	fmt.Println("Hello")
}

func foo(x, y int) int { // same type
	fmt.Println("Print in function:", x*y)
	return x + y
}

func foo2(x int) (int, int) {
	return x, x + 1
}

func cbv(y int) {
	y = y + 1
}

func cbr(y *int) {
	*y = *y + 1
}

func arr(x [3]int) int {
	return x[0]
}

func arrp(x *[3]int) {
	(*x)[0] = (*x)[0] + 1
}

func slif(sli []int) {
	sli[0] = sli[0] + 1
}

func main() {
	// Functions
	PrintHello()
	z := foo(2, 3)
	fmt.Println("Return Value:", z)

	a, b := foo2(4)
	fmt.Println(a, b)

	// Call By Value
	x := 2
	cbv(x)
	fmt.Println(x)

	// Call By Reference --> pass pointer
	cbr(&x)
	fmt.Println(x)

	// Passing Array Arguments --> can be big, so can be a problem
	ar := [3]int{1, 2, 3}
	fmt.Println(arr(ar))

	arrp(&ar)
	fmt.Println(ar) // messy and unnecessary --> use slice

	// Slice passes a pointer
	s := []int{1, 2, 3}
	slif(s)
	fmt.Println(s)
}
