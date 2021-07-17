package main

import (
	"fmt"
	"main/data"
	"math"
)

// type
type MyInt int

func (mi MyInt) Double() int { // receiver type --> implicit argument
	return int(mi * 2)
}

// struct
type Point struct {
	x float64
	y float64
}

func (p Point) DistToOrig() float64 {
	t := math.Pow(p.x, 2) + math.Pow(p.y, 2)
	return math.Sqrt(t)
}

func main() {
	v := MyInt(3)
	fmt.Println(v.Double())

	p1 := Point{5, 12}
	fmt.Println(p1.DistToOrig())

	data.PrintX()

	var p data.Point
	// No nee to reference --> p.Func()
	p.InitMe(3, 4)
	p.Scale(2)
	p.PrintMe()
	p.OffsetX(5)
	p.PrintMe()
}
