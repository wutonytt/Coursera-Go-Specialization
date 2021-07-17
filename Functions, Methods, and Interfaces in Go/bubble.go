package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Swap(slice []int, i int) {
	tmp := slice[i]
	slice[i] = slice[i+1]
	slice[i+1] = tmp
}

func BubbleSort(slice []int) {
	for i := 0; i < len(slice)-1; i++ {
		swapped := false
		for j := 0; j < len(slice)-i-1; j++ {
			if slice[j] > slice[j+1] {
				Swap(slice, j)
				swapped = true
			}
		}
		if swapped == false {
			break
		}
	}
}

func main() {
	var inputStr string
	var slice []int
	fmt.Println("Enter up to 10 integers seperated by ", ", e.g. 5,1,7")
	fmt.Scan(&inputStr)
	inputList := strings.Split(inputStr, ",")
	for _, elem := range inputList {
		tmp, _ := strconv.Atoi(elem)
		slice = append(slice, tmp)
	}
	BubbleSort(slice)
	for _, elem := range slice {
		fmt.Print(elem, " ")
	}
	fmt.Print("\n")
}
