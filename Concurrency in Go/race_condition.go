package main

import "fmt"

func main() {
	a := 0
	for i := 0; i < 10; i++ {
		go func() {
			a++
		}()
	}
	fmt.Println(a)

}

/*
	A race condition is when two or more routines
	have access to the same resource, such as a
	variable or data structure and attempt to read
	and write to that resource without any regard
	to the other routines.
*/
