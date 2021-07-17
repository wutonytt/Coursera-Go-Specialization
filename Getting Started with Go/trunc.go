package main

import "fmt"

func main() {
	fmt.Printf("Enter a floating point number: ")
	var input_float float32
	fmt.Scan(&input_float)
	fmt.Println(int(input_float))
}
