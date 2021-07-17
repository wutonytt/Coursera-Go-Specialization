package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	inputReader := bufio.NewReader(os.Stdin)

	fmt.Printf("Enter a name: ")
	inputName, _ := inputReader.ReadString('\n')
	inputName = strings.TrimSpace(inputName)

	fmt.Printf("Enter an address: ")
	inputAddr, _ := inputReader.ReadString('\n')
	inputAddr = strings.TrimSpace(inputAddr)

	nameAddrMap := map[string]string{
		"name":    inputName,
		"address": inputAddr,
	}
	barr, _ := json.Marshal(nameAddrMap)
	fmt.Println(string(barr))
}
