package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// pointer
	var x int = 1
	var y int
	var ip *int // pointer
	ip = &x
	y = *ip
	fmt.Println(y)

	// pointer: new
	ptr := new(int) // var ptr *int = new(int)
	*ptr = 3
	fmt.Println(ptr)

	s := "Joe"
	fmt.Println("Hello" + s)
	fmt.Printf("Hello %s\n", s)

	str := "Hello, I'm Tony."
	fmt.Println(strings.Replace(str, "l", "a", -1))

	conv := "123"
	fmt.Println(strconv.Atoi(conv))
	fmt.Println(strconv.Itoa(123))
	convf := "123.45"
	fmt.Println(strconv.ParseFloat(convf, 32))
	fmt.Println(strconv.FormatFloat(123.45, 'f', 4, 32))

	const (
		q = 1.5
		r = 2
	)
	fmt.Println(q, r)

	x = 6
	if x > 5 {
		fmt.Println(x - 1)
	}

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}

	// switch / case
	x = 2
	switch x {
	case 1:
		fmt.Println("case 1")
	case 2:
		fmt.Println("case 2")
	default:
		fmt.Println("nocase")

	}

	switch {
	case x > 1:
		fmt.Println("case 1")
	case x < 1:
		fmt.Println("case 2")
	default:
		fmt.Println("nocase")
	}

	// break / continue
	i = 0
	for i < 10 {
		i++
		if i == 5 {
			break
		}
		fmt.Println(i)
	}

	i = 0
	for i < 10 {
		i++
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}

	// scan
	fmt.Printf("Number of apples? ")
	var appleNum int
	// num, err := fmt.Scan(&appleNum)
	fmt.Scan(&appleNum)
	fmt.Println(appleNum)
}
