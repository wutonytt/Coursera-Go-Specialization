package main

import (
	"fmt"
)

func main() {
	// Arrays (fixed length)
	var x [5]int
	x[0] = 2
	fmt.Println(x[0], x[1])

	var y [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println(y[0], y[1])
	z := [...]int{6, 7, 8, 9}
	fmt.Println(z[0], z[1])
	fmt.Println(z)

	// Iterating Through Arrays
	for i, v := range z { // i: index; v: value
		fmt.Printf("ind %d, val %d\n", i, v)
	}

	// Slices (variable size)
	arr := [...]string{"a", "b", "c", "d", "e", "f", "g"}
	s1 := arr[1:3]
	s2 := arr[2:5]
	fmt.Println(s1)
	fmt.Println(s2)

	fmt.Println(len(s1), cap(s1))

	s1[1] = "x"
	fmt.Println(arr, s2[0])

	sli := []int{1, 2, 3} // it's a slice
	fmt.Println(sli)

	// Make an array/slice
	sli = make([]int, 10)
	fmt.Println(sli)

	sli = make([]int, 10, 15)
	fmt.Println(sli)

	// Append
	sli = make([]int, 3, 3)
	sli = append(sli, 100)
	fmt.Println(sli[:4])

	// Hash Table / Maps
	// var idMap map[string]int
	// idMap = make(map[string]int)

	idMap := map[string]int{
		"joe": 123,
	}
	fmt.Println(idMap["joe"])
	idMap["jane"] = 456
	fmt.Println(idMap["jane"])
	idMap["jane"] = 567
	fmt.Println(idMap["jane"])
	// delete(idMap, "joe")
	fmt.Println(idMap)

	id, p := idMap["jane"]
	fmt.Println(id, p)

	fmt.Println(len(idMap))

	for key, val := range idMap {
		fmt.Println(key, val)
	}

	// Struct
	type Person struct {
		name  string
		addr  string
		phone string
	}
	var p1 Person
	p1.name = "joe"
	p1.addr = "Taipei"
	p1.phone = "091029"
	ad := p1.addr
	fmt.Println(p1, ad)

	p2 := new(Person)
	fmt.Println(p2)

	p3 := Person{name: "jane", addr: "agea", phone: "0924"}
	fmt.Println(p3)
}
