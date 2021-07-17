package main

import (
	"fmt"
	"math"
	"time"
)

func GenDisplaceFn(a, v0, s0 float64) func(float64) float64 {
	fn := func(t float64) float64 {
		time.Sleep(time.Duration(t) * time.Second)
		s := 0.5*a*math.Pow(t, 2) + v0*t + s0
		return s
	}
	return fn
}

func main() {
	var a, v0, s0, t float64
	fmt.Printf("Enter a value for acceleration: ")
	fmt.Scan(&a)
	fmt.Printf("Enter a value for initial velocity: ")
	fmt.Scan(&v0)
	fmt.Printf("Enter a value for initial displacement: ")
	fmt.Scan(&s0)
	fmt.Printf("Enter a value for time: ")
	fmt.Scan(&t)

	// fmt.Printf("a: %f, v: %f, d: %f", a, v0, s0)
	fn := GenDisplaceFn(a, v0, s0)
	fmt.Println(fn(t))
}
