package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println(a * b / euclid(a, b))
}

func euclid(a, b int) int {
	if a < b {
		a, b = b, a
	}
	if a%b == 0 {
		return b
	}
	return euclid(b, a%b)
}
