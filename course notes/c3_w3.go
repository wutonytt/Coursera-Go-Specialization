package main

import (
	"fmt"
	"sync"
)

func foo1(wg *sync.WaitGroup) {
	fmt.Println("New Routine")
	wg.Done()
}

func prod(x, y int, c chan int) {
	c <- x * y
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go foo1(&wg)
	wg.Wait()
	fmt.Println("Main Routine")

	c := make(chan int)
	go prod(1, 2, c)
	go prod(3, 4, c)
	a := <-c
	b := <-c
	fmt.Println(a * b)
}
