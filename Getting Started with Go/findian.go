package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Printf("Enter a string: ")
	// var input_str string
	// fmt.Scanln(&input_str)

	inputReader := bufio.NewReader(os.Stdin)
	input_str, _ := inputReader.ReadString('\n')
	input_str = strings.TrimSpace(input_str)
	str := strings.ToLower(input_str)

	if strings.HasPrefix(str, "i") && strings.HasSuffix(str, "n") && strings.Contains(str, "a") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
