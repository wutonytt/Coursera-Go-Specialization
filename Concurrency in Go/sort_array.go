package main

import (
	"fmt"
	"sort"
)

func sortArray(arr []int, c chan []int) {
	sort.Ints(arr)
	c <- arr
}

func mergeArray(arr1, arr2 []int) []int {
	// fmt.Println("----Start mergeArray----")
	len1, len2 := len(arr1), len(arr2)

	idx1, idx2 := 0, 0

	var result []int

	for idx1 < len1 && idx2 < len2 {
		if arr1[idx1] <= arr2[idx2] {
			result = append(result, arr1[idx1])
			idx1++
		} else {
			result = append(result, arr2[idx2])
			idx2++
		}
	}
	if idx1 == len1 && idx2 < len2 {
		result = append(result, arr2[idx2:]...)
	} else if idx1 < len1 && idx2 == len2 {
		result = append(result, arr1[idx1:]...)
	}

	// fmt.Println(result)
	// fmt.Println("-----End mergeArray-----")
	return result
}

func main() {
	var length int
	var arr []int
	fmt.Println("Enter the length of the array of integers: ")
	fmt.Printf("> ")
	fmt.Scan(&length)
	fmt.Println("Enter the elements of the array of integers seperated by space: ")
	fmt.Printf("> ")
	for i := 0; i < length; i++ {
		var tmp int
		fmt.Scan(&tmp)
		arr = append(arr, tmp)
	}

	c := make(chan []int)
	step := length / 4

	go sortArray(arr[:step], c)
	go sortArray(arr[step:2*step], c)
	go sortArray(arr[2*step:3*step], c)
	go sortArray(arr[3*step:], c)

	for i := 0; i < length; i += step {
		go sortArray(arr[i:i+step], c)
	}

	// var result []int
	sub1 := <-c
	sub2 := <-c
	sub3 := <-c
	sub4 := <-c

	sub5 := mergeArray(sub1, sub2)
	sub6 := mergeArray(sub3, sub4)

	result := mergeArray(sub5, sub6)

	fmt.Println(result)
}
