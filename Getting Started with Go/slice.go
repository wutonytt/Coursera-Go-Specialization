package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var sli []int
	sli = make([]int, 0, 3)

	for {
		var input string
		fmt.Printf("Enter an integer: ")
		fmt.Scan(&input)
		// fmt.Println(input)
		if input == "X" {
			break
		}
		conv, _ := strconv.Atoi(input)
		sli = append(sli, conv)
		sort.Ints(sli)
		fmt.Println(sli)
	}
}
