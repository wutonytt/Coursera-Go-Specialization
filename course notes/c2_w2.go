package main

import (
	"fmt"
	"math"
)

// Variables as Functions
var funcVar func(int) int

func incFn(x int) int {
	return x + 1
}

// Functions as Arguments
func applyIt(afunct func(int) int, val int) int {
	return afunct(val)
}

// Function Defines a Function
func MakeDistOrigin(o_x, o_y float64) func(float64, float64) float64 {
	fn := func(x, y float64) float64 {
		return math.Sqrt(math.Pow(x-o_x, 2) + math.Pow(y-o_y, 2))
	}
	return fn
}

// Variable Argument Number (treated as a slice)
func getMax(vals ...int) int {
	maxVal := -1
	for _, v := range vals {
		if v > maxVal {
			maxVal = v
		}
	}
	return maxVal
}

func main() {
	// Variables as Functions
	funcVar = incFn
	fmt.Println(funcVar(1))

	// Functions as Arguments
	fmt.Println(applyIt(incFn, 2))

	// Anonymous Functions
	v := applyIt(func(x int) int { return x + 1 }, 3)
	fmt.Println(v)

	// Function Defines a Function
	Dist1 := MakeDistOrigin(0, 0)
	Dist2 := MakeDistOrigin(2, 2)
	fmt.Println(Dist1(2, 2))
	fmt.Println(Dist2(2, 2))

	// Variadic Slice Argument
	fmt.Println(getMax(1, 23, 2, 19))

	vslice := []int{1, 23, 2, 19}
	fmt.Println(getMax(vslice...)) // ...

	// Deferred Function Calls
	defer fmt.Println("Bye!")

	i := 1
	defer fmt.Println(i + 1) // arguments are evaluated immediately
	i++
	fmt.Println("Hello!")
}
