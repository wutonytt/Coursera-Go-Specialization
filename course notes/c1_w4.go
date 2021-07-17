package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
)

type Person struct {
	Name  string `json:"name"`
	Addr  string `json:"addr"`
	Phone string `json:"phone"`
}

func main() {
	http.Get("www.uci.edu")
	net.Dial("tcp", "uci.edu:80")

	// JSON
	p1 := Person{Name: "jane", Addr: "agea", Phone: "0924"}
	// fmt.Println(p1)
	barr, _ := json.MarshalIndent(p1, "", "	")
	fmt.Println(string(barr))

	var p2 Person
	json.Unmarshal(barr, &p2)
	fmt.Println(p2)

	// File Access

	// ioutil
	data, _ := ioutil.ReadFile("test.txt") // []byte
	fmt.Println(data)
	fmt.Println(string(data))

	dat := "Hello, world!"
	ioutil.WriteFile("outfile.txt", []byte(dat), 0777) // (filename, data, permission)

	// os
	f, _ := os.Open("test.txt")
	farr := make([]byte, 10)
	nb, _ := f.Read(farr)
	f.Close()
	fmt.Println(nb)
	fmt.Println(farr)
	fmt.Println(string(farr))

	fo, _ := os.Create("outfile1.txt")
	fostr := "OUT Test"
	foarr := []byte{73, 101, 114}
	nbo, _ := fo.Write(foarr)
	fo.WriteString(fostr)
	fmt.Println(nbo)
}
