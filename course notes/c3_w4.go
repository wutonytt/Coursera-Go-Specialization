package main

import (
	"fmt"
	"sync"
)

var i int = 0
var mut sync.Mutex
var wg sync.WaitGroup
var on sync.Once

func setup() {
	fmt.Println("Init")
}

func inc() {
	on.Do(setup)
	mut.Lock()
	i++
	mut.Unlock()
	fmt.Println("Hello")
	wg.Done()
}

func main() {
	wg.Add(2)
	go inc()
	go inc()
	wg.Wait()
}
