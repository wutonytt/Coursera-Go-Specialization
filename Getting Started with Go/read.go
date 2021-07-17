package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Name struct {
	fname string
	lname string
}

func main() {
	var filename string
	fmt.Printf("Enter the name of the text file: ")
	fmt.Scan(&filename)

	file, _ := os.Open(filename)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var sli []Name

	for scanner.Scan() {
		text := scanner.Text()
		splitText := strings.Split(text, " ")
		n := Name{fname: splitText[0], lname: splitText[1]}
		sli = append(sli, n)
	}
	file.Close()

	for _, x := range sli {
		fmt.Println(x.fname, x.lname)
	}
}
