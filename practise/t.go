package main

import "fmt"

var t interface{}

func main() {

	m := make(map[string]string)
	m["a"] = "1"
	m["b"] = "2"
	for _, v := range m {
		fmt.Println("value:", v)

	}

}
