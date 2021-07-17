package main

type Shape2D interface {
	Area() float64
	Perimeter() float64
}

type Triangle {}
func (t Triangle) Area() float64 {}
func (t Triangle) Perimeter() float64 {}

type Rectangle {}
func (t Rectangle) Area() float64 {}
func (t Rectangle) Perimeter() float64 {}

func FitInYard(s Shape2D) bool {
	if (s.Area() < 100 && s.Perimeter()) {
		return true
	}
	return false
}



func main() {

}