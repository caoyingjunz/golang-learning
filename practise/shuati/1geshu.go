package main

import (
	"fmt"
)

func solve(n int) int {
	var count int
	for n > 0 {
		if n&1 == 1 {
			count++
		}
		n = n >> 1
	}
	return count
}

func main() {
	var n int
	for {
		_, err := fmt.Scanln(&n)
		if err != nil {
			break
		}
		fmt.Println(solve(n))
	}
}
